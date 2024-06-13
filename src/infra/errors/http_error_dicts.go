package errors

import (
	"net/http"
)

var httpCode = map[ErrorCode]int{
	UNKNOWN_ERROR:              http.StatusInternalServerError,
	DATA_INVALID:               http.StatusBadRequest,
	STATUS_PAGE_NOT_FOUND:      http.StatusNotFound,
	UNAUTHORIZED:               http.StatusUnauthorized,
	FAILED_RETRIEVE_DATA:       http.StatusInternalServerError,
	USER_ALREADY_EXIST:         http.StatusConflict,
	INVALID_AMOUNT:             http.StatusBadRequest,
	DESTINATION_USER_NOT_FOUND: http.StatusNotFound,
	INSUFICIENT_BALANCE:        http.StatusBadRequest,
	TO_OWN_ACCOUNT:             http.StatusBadRequest,
}
