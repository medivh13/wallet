package transaction

import (
	"context"
	"encoding/json"
	"log"
	"time"
	dto "wallet/src/app/dto/transaction"

	repo "wallet/src/infra/persistence/postgres/transaction"
	redis "wallet/src/infra/persistence/redis/service"
)

type TransactionUCInterface interface {
	Transfer(data *dto.TransferReqDTO) error
	GetTopTransactionByUser(walletID int64) ([]*dto.GetTopTransRespDTO, error)
	GetOverallTopTransactions(ctx context.Context) ([]*dto.GetOverallRespDTO, error)
}

type transactionUseCase struct {
	Repo  repo.TransactionRepository
	Redis redis.ServRedisInt
}

func NewTransactionUseCase(repo repo.TransactionRepository, redis redis.ServRedisInt) TransactionUCInterface {
	return &transactionUseCase{
		Repo:  repo,
		Redis: redis,
	}
}

func (uc *transactionUseCase) Transfer(data *dto.TransferReqDTO) error {
	err := uc.Repo.Transfer(data)
	if err != nil {
		return err
	}

	return nil
}

func (uc *transactionUseCase) GetTopTransactionByUser(walletID int64) ([]*dto.GetTopTransRespDTO, error) {
	resp, err := uc.Repo.GetTopTransactionByUser(walletID)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (uc *transactionUseCase) GetOverallTopTransactions(ctx context.Context) ([]*dto.GetOverallRespDTO, error) {
	const cacheKeyTopUsers = "top_10_users_by_debit_value"
	var resp []*dto.GetOverallRespDTO

	dataRedis, err := uc.Redis.GetData(ctx, cacheKeyTopUsers)

	if err != nil {
		log.Printf("unable to GET data from redis. error: %v", err)
	}

	if dataRedis != "" {
		_ = json.Unmarshal([]byte(dataRedis), &resp)

		log.Println("data from redis")
		return resp, nil

	}

	resp, err = uc.Repo.GetOverallTopTransactions()
	if err != nil {
		return resp, err
	}

	redisData, _ := json.Marshal(resp)
	ttl := time.Duration(2) * time.Minute

	err = uc.Redis.SetData(context.Background(), cacheKeyTopUsers, redisData, ttl)
	if err != nil {
		log.Printf("unable to SET data. error: %v", err)
		return nil, err
	}

	return resp, nil
}
