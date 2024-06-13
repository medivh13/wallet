package usecases

import (
	balanceUC "wallet/src/app/usecases/balance"
	transactionUC "wallet/src/app/usecases/transaction"
	userUC "wallet/src/app/usecases/user"
)

type AllUseCases struct {
	UserUC        userUC.UserUCInterface
	BalanceUC     balanceUC.BalanceUCInterface
	TransactionUC transactionUC.TransactionUCInterface
}
