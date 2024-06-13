package balance

import (
	"errors"
	"testing"
	mockDTO "wallet/mocks/app/dto/balance"

	mockRepo "wallet/mocks/infra/persistence/postgres/balance"
	dto "wallet/src/app/dto/balance"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockBalanceUseCase struct {
	mock.Mock
}

type BalanceUseCaseList struct {
	suite.Suite
	mockDTO        *mockDTO.MockBalanceDTO
	mockRepo       *mockRepo.MockBalanceRepo
	useCase        BalanceUCInterface
	dtoTopUpReq    *dto.TopUpReqDTO
	dtoBalanceReq  *dto.BalanceReqDTO
	dtoBalanceResp *dto.BalanceRespDTO
}

func (suite *BalanceUseCaseList) SetupTest() {
	suite.mockDTO = new(mockDTO.MockBalanceDTO)
	suite.mockRepo = new(mockRepo.MockBalanceRepo)
	suite.useCase = NewBalanceUseCase(suite.mockRepo)

	suite.dtoTopUpReq = &dto.TopUpReqDTO{
		Amount:   50000,
		WalletID: 3,
	}

	suite.dtoBalanceReq = &dto.BalanceReqDTO{
		WalletID: 3,
	}

	suite.dtoBalanceResp = &dto.BalanceRespDTO{
		Balance: 100000,
	}
}

func (u *BalanceUseCaseList) TestTopUpSuccess() {
	u.mockRepo.Mock.On("TopUp", u.dtoTopUpReq).Return(nil)
	err := u.useCase.TopUp(u.dtoTopUpReq)
	u.Equal(nil, err)
}

func (u *BalanceUseCaseList) TestTopUpFailed() {
	u.mockRepo.Mock.On("TopUp", u.dtoTopUpReq).Return(errors.New(mock.Anything))
	err := u.useCase.TopUp(u.dtoTopUpReq)
	u.Equal(errors.New(mock.Anything), err)
}

func (u *BalanceUseCaseList) TestGetSuccess() {
	u.mockRepo.Mock.On("Get", u.dtoBalanceReq).Return(mock.Anything, nil)
	_, err := u.useCase.Get(u.dtoBalanceReq)
	u.Equal(nil, err)
}

func (u *BalanceUseCaseList) TestGetFailed() {
	u.mockRepo.Mock.On("Get", u.dtoBalanceReq).Return(nil, errors.New(mock.Anything))
	_, err := u.useCase.Get(u.dtoBalanceReq)
	u.Equal(errors.New(mock.Anything), err)
}

func TestUsecase(t *testing.T) {
	suite.Run(t, new(BalanceUseCaseList))
}
