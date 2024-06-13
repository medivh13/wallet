package mock_transaction

import (
	dto "wallet/src/app/dto/transaction"

	"github.com/stretchr/testify/mock"
)

type MockTransactionDTO struct {
	mock.Mock
}

func NewMockTransactionDTO() *MockTransactionDTO {
	return &MockTransactionDTO{}
}

var _ dto.TransactionDTOInterface = &MockTransactionDTO{}

func (m *MockTransactionDTO) Validate() error {
	args := m.Called()
	var err error
	if n, ok := args.Get(0).(error); ok {
		err = n
		return err
	}

	return nil
}
