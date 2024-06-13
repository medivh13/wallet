package mock_user

import (
	dto "wallet/src/app/dto/user"
	repo "wallet/src/infra/persistence/postgres/user"

	"github.com/stretchr/testify/mock"
)

type MockUsersRepo struct {
	mock.Mock
}

func NewMockUsersRepo() *MockUsersRepo {
	return &MockUsersRepo{}
}

var _ repo.UserRepository = &MockUsersRepo{}

func (m *MockUsersRepo) Register(data *dto.RegisterReqDTO) (*dto.RegisterRespDTO, error) {
	args := m.Called(data)
	var (
		res *dto.RegisterRespDTO
		err error
	)

	if n, ok := args.Get(0).(*dto.RegisterRespDTO); ok {

		res = n
	}

	if n, ok := args.Get(1).(error); ok {

		err = n
	}

	return res, err
}

func (m *MockUsersRepo) Login(data *dto.LoginReqDTO) (*dto.RegisterRespDTO, error) {
	args := m.Called(data)
	var (
		res *dto.RegisterRespDTO
		err error
	)

	if n, ok := args.Get(0).(*dto.RegisterRespDTO); ok {

		res = n
	}

	if n, ok := args.Get(1).(error); ok {

		err = n
	}

	return res, err
}
