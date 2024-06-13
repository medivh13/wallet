package balance

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	dto "wallet/src/app/dto/balance"
	usecases "wallet/src/app/usecases/balance"
	common_error "wallet/src/infra/errors"
	"wallet/src/infra/helper"
	"wallet/src/interface/rest/response"
)

type BalanceHandlerInterface interface {
	TopUp(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type balanceHandler struct {
	response response.IResponseClient
	usecase  usecases.BalanceUCInterface
}

func NewBalanceHandler(r response.IResponseClient, h usecases.BalanceUCInterface) BalanceHandlerInterface {
	return &balanceHandler{
		response: r,
		usecase:  h,
	}
}

func (h *balanceHandler) TopUp(w http.ResponseWriter, r *http.Request) {
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

	postDTO := dto.TopUpReqDTO{}
	err = json.NewDecoder(r.Body).Decode(&postDTO)
	if err != nil {
		log.Println(err)
		h.response.HttpError(w, common_error.NewError(common_error.INVALID_AMOUNT, err))
		return
	}
	err = postDTO.Validate()
	postDTO.WalletID = dataClaim.WalletID

	if err != nil {
		log.Println(err)
		h.response.HttpError(w, common_error.NewError(common_error.INVALID_AMOUNT, err))
		return
	}

	err = h.usecase.TopUp(&postDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.FAILED_CREATE_DATA, err))
		return
	}
	w.WriteHeader(http.StatusNoContent)
	h.response.JSON(
		w,
		"TopUp Successful",
		nil,
		nil,
	)
}

func (h *balanceHandler) Get(w http.ResponseWriter, r *http.Request) {
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

	getDTO := dto.BalanceReqDTO{}

	getDTO.WalletID = dataClaim.WalletID

	data, err := h.usecase.Get(&getDTO)
	if err != nil {
		h.response.HttpError(w, common_error.NewError(common_error.UNKNOWN_ERROR, err))
		return
	}

	h.response.JSON(
		w,
		"Balance Read Success",
		data,
		nil,
	)
}
