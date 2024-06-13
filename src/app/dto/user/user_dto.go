package user

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type UserReqDTOInterface interface {
	Validate() error
}

type RegisterReqDTO struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func (dto *RegisterReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.UserName, validation.Required),
		validation.Field(&dto.Password, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type RegisterRespDTO struct {
	Token    string `json:"token"`
}

type LoginReqDTO struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func (dto *LoginReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.UserName, validation.Required),
		validation.Field(&dto.Password, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type RegisterModel struct {
	ID       int64  `db:"id"`
	UserName string `db:"username"`
	Password string `db:"password"`
	WalletID int64  `db:"wallet_id"`
}
