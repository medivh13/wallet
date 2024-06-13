package errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetKnownHttpStatus(t *testing.T) {
	httpCode[1] = http.StatusBadRequest
	commonError := NewError(1, nil)

	assert.Equal(t, http.StatusBadRequest, commonError.GetHttpStatus(), "It will return HTTP status 400")
}

func TestGetUnknownHttpStatus(t *testing.T) {
	delete(httpCode, 1)
	commonError := NewError(1, nil)

	assert.Equal(t, http.StatusInternalServerError, commonError.GetHttpStatus(), "It will return HTTP status 500")
}

func TestHttpErrorString(t *testing.T) {
	errorCodes[1] = &CommonError{
		ClientMessage: "this is a client message",
	}

	commonError := NewError(1, nil)

	assert.Equal(t, "this is a client message", commonError.ToHttpError().Error(), "It will return string \"this is a client message\"")
}

func TestHttpErrorHasCommonError(t *testing.T) {
	commonError := NewError(0, errors.New("this is an error"))

	assert.NotNil(t, commonError.ClientMessage, "HttpError should have ClientMessage")
	assert.NotNil(t, commonError.SystemMessage, "HttpError should have SystemMessage")
	assert.NotNil(t, commonError.ErrorCode, "HttpError should have ErrorCode")
	assert.NotNil(t, commonError.ErrorMessage, "HttpError should have ErrorMessage")
	assert.NotNil(t, commonError.ErrorTrace, "HttpError should have ErrorTrace")
}
