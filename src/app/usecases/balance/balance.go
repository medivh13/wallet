package balance

import (
	"log"

	dto "wallet/src/app/dto/balance"

	repo "wallet/src/infra/persistence/postgres/balance"
)

type BalanceUCInterface interface {
	TopUp(data *dto.TopUpReqDTO) error
	Get(data *dto.BalanceReqDTO) (*dto.BalanceRespDTO, error)
}

type balanceUseCase struct {
	Repo repo.BalanceRepository
}

func NewBalanceUseCase(repo repo.BalanceRepository) BalanceUCInterface {
	return &balanceUseCase{
		Repo: repo,
	}
}

func (uc *balanceUseCase) TopUp(data *dto.TopUpReqDTO) error {
	err := uc.Repo.TopUp(data)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (uc *balanceUseCase) Get(data *dto.BalanceReqDTO) (*dto.BalanceRespDTO, error) {
	resp, err := uc.Repo.Get(data)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return resp, nil
}
