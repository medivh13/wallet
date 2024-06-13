package postgres

/*
 * Author      : Jody (jody.almaida@gmail)
 * Modifier    :
 * Domain      : wallet
 */

import (
	"fmt"

	"wallet/src/infra/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type PostgresDb struct {
	Conn *sqlx.DB
}

func New(conf config.SqlDbConf, logger *logrus.Logger) (PostgresDb, error) {
	db := PostgresDb{}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		conf.Host,
		conf.Username,
		conf.Password,
		conf.Name,
		conf.Port,
	)

	if conf.Password == "" {
		dsn = fmt.Sprintf(
			"host=%s user=%s dbname=%s port=%s sslmode=disable",
			conf.Host,
			conf.Username,
			conf.Name,
			conf.Port,
		)
	}

	conn, err := sqlx.Open("postgres", dsn)
	if err != nil {
		panic("Failed to connect to database!")
	}

	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}

	logger.Printf("sql database connection %s success", db.Conn.DriverName())
	return db, nil
}
