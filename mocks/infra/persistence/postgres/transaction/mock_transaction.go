package mock_transaction

import (
	dto "wallet/src/app/dto/transaction"
	repo "wallet/src/infra/persistence/postgres/transaction"

	"github.com/stretchr/testify/mock"
)

type MockTransactionRepo struct {
	mock.Mock
}

func NewMockTransactionRepo() *MockTransactionRepo {
	return &MockTransactionRepo{}
}

var _ repo.TransactionRepository = &MockTransactionRepo{}

func (m *MockTransactionRepo) Transfer(data *dto.TransferReqDTO) error {
	args := m.Called(data)
	var (
		err error
	)

	if n, ok := args.Get(0).(error); ok {

		err = n
	}

	return err
}

func (m *MockTransactionRepo) GetTopTransactionByUser(walletID int64) ([]*dto.GetTopTransRespDTO, error) {
	args := m.Called(walletID)
	var (
		res []*dto.GetTopTransRespDTO
		err error
	)

	if n, ok := args.Get(0).([]*dto.GetTopTransRespDTO); ok {

		res = n
	}

	if n, ok := args.Get(1).(error); ok {

		err = n
	}

	return res, err
}

func (m *MockTransactionRepo) GetOverallTopTransactions() ([]*dto.GetOverallRespDTO, error) {
	args := m.Called()
	var (
		res []*dto.GetOverallRespDTO
		err error
	)

	if n, ok := args.Get(0).([]*dto.GetOverallRespDTO); ok {

		res = n
	}

	if n, ok := args.Get(1).(error); ok {

		err = n
	}

	return res, err
}
