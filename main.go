package main

import (
	"context"
	"database/sql"

	usecases "wallet/src/app/usecases"

	"wallet/src/infra/config"

	postgres "wallet/src/infra/persistence/postgres"
	"wallet/src/infra/persistence/redis"

	balanceRepo "wallet/src/infra/persistence/postgres/balance"
	transactionRepo "wallet/src/infra/persistence/postgres/transaction"
	userRepo "wallet/src/infra/persistence/postgres/user"

	redisServe "wallet/src/infra/persistence/redis/service"
	"wallet/src/interface/rest"

	ms_log "wallet/src/infra/log"

	balanceUC "wallet/src/app/usecases/balance"
	transactionUC "wallet/src/app/usecases/transaction"
	userUC "wallet/src/app/usecases/user"

	_ "github.com/joho/godotenv/autoload"

	"github.com/sirupsen/logrus"
)


func main() {
	
	ctx := context.Background()

	conf := config.Make()

	isProd := false
	if conf.App.Environment == "PRODUCTION" {
		isProd = true
	}

	m := make(map[string]interface{})
	m["env"] = conf.App.Environment
	m["service"] = conf.App.Name
	logger := ms_log.NewLogInstance(
		ms_log.LogName(conf.Log.Name),
		ms_log.IsProduction(isProd),
		ms_log.LogAdditionalFields(m))

	postgresdb, err := postgres.New(conf.SqlDb, logger)
	redisClient, err := redis.NewRedisClient(conf.Redis, logger)

	redisServe := redisServe.NewServRedis(redisClient)

	defer func(l *logrus.Logger, sqlDB *sql.DB, dbName string) {
		err := sqlDB.Close()
		if err != nil {
			l.Errorf("error closing sql database %s: %s", dbName, err)
		} else {
			l.Printf("sql database %s successfuly closed.", dbName)
		}
	}(logger, postgresdb.Conn.DB, postgresdb.Conn.DriverName())

	userRepository := userRepo.NewUserRepository(postgresdb.Conn)
	balanceRepository := balanceRepo.NewBalanceRepository(postgresdb.Conn)
	transactionRepository := transactionRepo.NewTransactionRepository(postgresdb.Conn)

	httpServer, err := rest.New(
		conf.Http,
		isProd,
		logger,
		usecases.AllUseCases{
			UserUC:        userUC.NewUserUseCase(userRepository),
			BalanceUC:     balanceUC.NewBalanceUseCase(balanceRepository),
			TransactionUC: transactionUC.NewTransactionUseCase(transactionRepository, redisServe),
		},
	)
	if err != nil {
		panic(err)
	}
	httpServer.Start(ctx)

}
