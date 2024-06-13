package transaction

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	dto "wallet/src/app/dto/transaction"
	usecases "wallet/src/app/usecases/transaction"
	common_error "wallet/src/infra/errors"
	"wallet/src/infra/helper"
	"wallet/src/interface/rest/response"
)

type TransactionHandlerInterface interface {
	Transfer(w http.ResponseWriter, r *http.Request)
	GetTopTen(w http.ResponseWriter, r *http.Request)
	GetOverallTopTransactions(w http.ResponseWriter, r *http.Request)
}

type transactionHandler struct {
	response response.IResponseClient
	usecase  usecases.TransactionUCInterface
}

func NewTransactionHandler(r response.IResponseClient, h usecases.TransactionUCInterface) TransactionHandlerInterface {
	return &transactionHandler{
		response: r,
		usecase:  h,
	}
}

func (h *transactionHandler) Transfer(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		h.response.HttpError(w, common_error.NewError(common_error.UNAUTHORIZED, errors.New("missing authorization token")))
		return
	}

	// Verifikasi token
	dataClaim, err := helper.VerifyToken(tokenString)
	if err != nil {
		log.Println(err)
		h.response.HttpError(w, common_error.NewError(common_error.UNAUTHORIZED, err))
		return
	}

	postDTO := dto.TransferReqDTO{}
	err = json.NewDecoder(r.Body).Decode(&postDTO)
	if err != nil {
		log.Println(err)
		h.response.HttpError(w, common_error.NewError(common_error.DATA_INVALID, err))
		return
	}
	err = postDTO.Validate()
	postDTO.WalletIDSender = dataClaim.WalletID

	if err != nil {
		log.Println(err)
		h.response.HttpError(w, common_error.NewError(common_error.INVALID_AMOUNT, err))
		return
	}

	if postDTO.ToUserName == dataClaim.UserName {
		err := errors.New("you cannot transfer to your own account")
		h.response.HttpError(w, common_error.NewError(common_error.TO_OWN_ACCOUNT, err))
		return
	}

	err = h.usecase.Transfer(&postDTO)
	if err != nil {

		if errors.Is(err, helper.ErrInsufficientBalance) {
			h.response.HttpError(w, common_error.NewError(common_error.INSUFICIENT_BALANCE, err))
			return
		}
		if errors.Is(err, helper.ErrUserNotFound) {
			h.response.HttpError(w, common_error.NewError(common_error.DESTINATION_USER_NOT_FOUND, err))
			return
		}
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_CREATE_DATA, err))
		return
	}
	w.WriteHeader(http.StatusNoContent)
	h.response.JSON(
		w,
		"Transfer Success",
		nil,
		nil,
	)
}

func (h *transactionHandler) GetTopTen(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		h.response.HttpError(w, common_error.NewError(common_error.UNAUTHORIZED, errors.New("missing authorization token")))
		return
	}

	dataClaim, err := helper.VerifyToken(tokenString)
	if err != nil {
		log.Println(err)
		h.response.HttpError(w, common_error.NewError(common_error.UNAUTHORIZED, err))
		return
	}

	data, err := h.usecase.GetTopTen(dataClaim.WalletID)

	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_RETRIEVE_DATA, err))
		return
	}

	h.response.JSON(
		w,
		"Success",
		data,
		nil,
	)
}

func (h *transactionHandler) GetOverallTopTransactions(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		h.response.HttpError(w, common_error.NewError(common_error.UNAUTHORIZED, errors.New("missing authorization token")))
		return
	}

	_, err := helper.VerifyToken(tokenString)
	if err != nil {
		log.Println(err)
		h.response.HttpError(w, common_error.NewError(common_error.UNAUTHORIZED, err))
		return
	}

	data, err := h.usecase.GetOverallTopTransactions(r.Context())

	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_RETRIEVE_DATA, err))
		return
	}

	h.response.JSON(
		w,
		"Success",
		data,
		nil,
	)
}
