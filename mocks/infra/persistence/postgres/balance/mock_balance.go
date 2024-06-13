package mock_balance

import (
	dto "wallet/src/app/dto/balance"
	repo "wallet/src/infra/persistence/postgres/balance"

	"github.com/stretchr/testify/mock"
)

type MockBalanceRepo struct {
	mock.Mock
}

func NewMockBalanceRepo() *MockBalanceRepo {
	return &MockBalanceRepo{}
}

var _ repo.BalanceRepository = &MockBalanceRepo{}

func (m *MockBalanceRepo) TopUp(data *dto.TopUpReqDTO) error {
	args := m.Called(data)
	var (
		err error
	)

	if n, ok := args.Get(0).(error); ok {

		err = n
	}

	return err
}

func (m *MockBalanceRepo) Get(data *dto.BalanceReqDTO) (*dto.BalanceRespDTO, error) {
	args := m.Called(data)
	var (
		res *dto.BalanceRespDTO
		err error
	)

	if n, ok := args.Get(0).(*dto.BalanceRespDTO); ok {

		res = n
	}

	if n, ok := args.Get(1).(error); ok {

		err = n
	}

	return res, err
}
