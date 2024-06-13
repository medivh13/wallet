package mock_user

import (
	dto "wallet/src/app/dto/user"

	"github.com/stretchr/testify/mock"
)

type MockUsersDTO struct {
	mock.Mock
}

func NewMockUsersDTO() *MockUsersDTO {
	return &MockUsersDTO{}
}

var _ dto.UserReqDTOInterface = &MockUsersDTO{}

func (m *MockUsersDTO) Validate() error {
	args := m.Called()
	var err error
	if n, ok := args.Get(0).(error); ok {
		err = n
		return err
	}

	return nil
}
