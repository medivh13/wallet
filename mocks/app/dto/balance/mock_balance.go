package mock_balance

import (
	dto "wallet/src/app/dto/balance"

	"github.com/stretchr/testify/mock"
)

type MockBalanceDTO struct {
	mock.Mock
}

func NewMockBalanceDTO() *MockBalanceDTO {
	return &MockBalanceDTO{}
}

var _ dto.BalanceDTOInterface = &MockBalanceDTO{}

func (m *MockBalanceDTO) Validate() error {
	args := m.Called()
	var err error
	if n, ok := args.Get(0).(error); ok {
		err = n
		return err
	}

	return nil
}
