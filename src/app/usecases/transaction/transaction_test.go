package transaction

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"
	mockDTO "wallet/mocks/app/dto/transaction"

	mockRepo "wallet/mocks/infra/persistence/postgres/transaction"
	mockRedis "wallet/mocks/infra/persistence/redis"
	dto "wallet/src/app/dto/transaction"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockTransactionUseCase struct {
	mock.Mock
}

type TransactionUseCaseList struct {
	suite.Suite
	mockDTO           *mockDTO.MockTransactionDTO
	mockRepo          *mockRepo.MockTransactionRepo
	useCase           TransactionUCInterface
	dtoTransferReqDTO *dto.TransferReqDTO
	dtoGetOveralResp  []*dto.GetOverallRespDTO
	mockRedis         *mockRedis.MockRedis
	key               string
}

func (suite *TransactionUseCaseList) SetupTest() {
	suite.mockDTO = new(mockDTO.MockTransactionDTO)
	suite.mockRepo = new(mockRepo.MockTransactionRepo)
	suite.mockRedis = new(mockRedis.MockRedis)
	suite.useCase = NewTransactionUseCase(suite.mockRepo, suite.mockRedis)

	suite.dtoTransferReqDTO = &dto.TransferReqDTO{
		ToUserName:     "jody",
		Amount:         50000,
		WalletIDSender: 3,
	}

	suite.dtoGetOveralResp = []*dto.GetOverallRespDTO{
		{
			UserName: "jody",
			Amount:   100000,
		},
	}

	suite.key = "top_10_users_by_debit_value"

}

func (u *TransactionUseCaseList) TestTransferSuccess() {
	u.mockRepo.Mock.On("Transfer", u.dtoTransferReqDTO).Return(nil)
	err := u.useCase.Transfer(u.dtoTransferReqDTO)
	u.Equal(nil, err)
}

func (u *TransactionUseCaseList) TestTransferFailed() {
	u.mockRepo.Mock.On("Transfer", u.dtoTransferReqDTO).Return(errors.New(mock.Anything))
	err := u.useCase.Transfer(u.dtoTransferReqDTO)
	u.Equal(errors.New(mock.Anything), err)
}

func (u *TransactionUseCaseList) TestGetTopTenSuccess() {
	u.mockRepo.Mock.On("GetTopTransactionByUser", int64(3)).Return(mock.Anything, nil)
	_, err := u.useCase.GetTopTransactionByUser(int64(3))
	u.Equal(nil, err)
}

func (u *TransactionUseCaseList) TestGetTopTenFailed() {
	u.mockRepo.Mock.On("GetTopTransactionByUser", int64(3)).Return(nil, errors.New(mock.Anything))
	_, err := u.useCase.GetTopTransactionByUser(int64(3))
	u.Equal(errors.New(mock.Anything), err)
}

func (u *TransactionUseCaseList) TestGetOveralFromRedisSuccess() {
	ctx := context.Background()
	dataresp, _ := json.Marshal(u.dtoGetOveralResp)
	u.mockRedis.Mock.On("GetData", ctx, u.key).Return(string(dataresp), nil)
	_, err := u.useCase.GetOverallTopTransactions(ctx)
	u.Equal(nil, err)
}

func (u *TransactionUseCaseList) TestGetOveralSuccess() {
	ctx := context.Background()
	u.mockRedis.Mock.On("GetData", ctx, u.key).Return("", errors.New(mock.Anything))
	u.mockRepo.Mock.On("GetOverallTopTransactions").Return(mock.Anything, nil)
	u.mockRedis.Mock.On("SetData", ctx, u.key, mock.Anything, time.Duration(2)*time.Minute).Return(nil)
	_, err := u.useCase.GetOverallTopTransactions(ctx)
	u.Equal(nil, err)
}

func (u *TransactionUseCaseList) TestGetOveralFail() {
	ctx := context.Background()
	u.mockRedis.Mock.On("GetData", ctx, u.key).Return("", errors.New(mock.Anything))
	u.mockRepo.Mock.On("GetOverallTopTransactions").Return(nil, errors.New(mock.Anything))
	_, err := u.useCase.GetOverallTopTransactions(ctx)
	u.Equal(errors.New(mock.Anything), err)
}

	func (u *TransactionUseCaseList) TestGetOveralSetDataRedisFail() {
		ctx := context.Background()
		u.mockRedis.Mock.On("GetData", ctx, u.key).Return("", errors.New(mock.Anything))
		u.mockRepo.Mock.On("GetOverallTopTransactions").Return(mock.Anything, nil)
		u.mockRedis.Mock.On("SetData", ctx, u.key, mock.Anything, time.Duration(2)*time.Minute).Return(errors.New(mock.Anything))
		_, err := u.useCase.GetOverallTopTransactions(ctx)
		u.Equal(errors.New(mock.Anything), err)
	}
func TestUsecase(t *testing.T) {
	suite.Run(t, new(TransactionUseCaseList))
}
