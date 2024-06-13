package transaction

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type TransactionDTOInterface interface {
	Validate() error
}

type TransferReqDTO struct {
	ToUserName     string  `json:"to_username"`
	Amount         float64 `json:"amount"`
	WalletIDSender int64
}

func (dto *TransferReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.Amount,
			validation.Required,
			validation.Min(1.00),       // memastikan amount positif dan lebih dari 0
			validation.Max(9999999.99), // memastikan amount kurang dari 10,000,000
		),
		validation.Field(&dto.ToUserName, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type GetTopTenRespDTO struct {
	UserName string  `json:"username"`
	Amount   float64 `json:"amount"`
}

type GetOverallRespDTO struct {
	UserName string  `json:"username" db:"username"`
	Amount   float64 `json:"transacted_value" db:"transacted_value"`
}
