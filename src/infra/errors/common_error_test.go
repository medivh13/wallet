package errors

import (
	"errors"
	"testing"

	validation "github.com/go-ozzo/ozzo-validation"

	"github.com/stretchr/testify/assert"
)

func TestOnlyNewClientErrorMessage(t *testing.T) {
	errMsg := NewError(0, nil)
	errMsg.SetClientMessage("This is a new client message")

	if assert.NotNil(t, errMsg.ClientMessage) {
		assert.Equal(t, "This is a new client message", errMsg.ClientMessage, "Client message should be a new client message")
	}

	if assert.NotNil(t, errMsg.SystemMessage) {
		assert.Equal(t, errMsg.SystemMessage, "Unknown error.")
	}
}

func TestOnlyNewSystemErrorMessage(t *testing.T) {
	errMsg := NewError(0, nil)
	errMsg.SetSystemMessage("This is a new system message")

	if assert.NotNil(t, errMsg.ClientMessage) {
		assert.Equal(t, errMsg.ClientMessage, "Unknown error.")
	}

	if assert.NotNil(t, errMsg.SystemMessage) {
		assert.Equal(t, "This is a new system message", errMsg.SystemMessage, "System message should be a new system message")
	}
}

func TestNewErrorMessage(t *testing.T) {
	errMsg := NewError(0, nil)
	errMsg.SetClientMessage("This is a new client message")
	errMsg.SetSystemMessage("This is a new system message")

	if assert.NotNil(t, errMsg.ClientMessage) {
		assert.Equal(t, "This is a new client message", errMsg.ClientMessage, "Client message should be a new client message")
	}

	if assert.NotNil(t, errMsg.SystemMessage) {
		assert.Equal(t, "This is a new system message", errMsg.SystemMessage, "System message should be a new system message")
	}
}

func TestCommonError(t *testing.T) {
	errMsg := NewError(0, nil)

	if assert.NotNil(t, errMsg.ClientMessage) {
		assert.Equal(t, "Unknown error.", errMsg.ClientMessage, "Client message should be a new client message")
	}

	if assert.NotNil(t, errMsg.SystemMessage) {
		assert.Equal(t, "Unknown error.", errMsg.SystemMessage, "System message should be a new system message")
	}
}

func TestThrownError(t *testing.T) {
	errMsg := NewError(0, errors.New("this is another error"))

	if assert.NotNil(t, errMsg.ClientMessage) {
		assert.Equal(t, "Unknown error.", errMsg.ClientMessage, "Client message should be a new client message")
	}

	if assert.NotNil(t, errMsg.SystemMessage) {
		assert.Equal(t, "Unknown error.", errMsg.SystemMessage, "System message should be a new system message")
	}

	if assert.NotNil(t, errMsg.ErrorMessage) {
		assert.Equal(t, "this is another error", *errMsg.ErrorMessage, "System message should be a new system message")
	}
}

func TestReThrowCommonError(t *testing.T) {
	errorCodes[1] = &CommonError{
		ClientMessage: "This is my client message",
		SystemMessage: "This is my system message",
		ErrorCode:     1,
	}
	errorCodes[2] = &CommonError{
		ClientMessage: "This is my second client message",
		SystemMessage: "This is my second system message",
		ErrorCode:     2,
	}

	errMsg := NewError(1, errors.New("this is an error"))
	errMsg2 := NewError(2, errMsg)

	if assert.NotNil(t, errMsg2.ClientMessage) {
		assert.Equal(t, "This is my client message", errMsg2.ClientMessage, "Client message should be \"This is my client message\"")
	}

	if assert.NotNil(t, errMsg2.SystemMessage) {
		assert.Equal(t, "This is my system message", errMsg2.SystemMessage, "System message should be \"This is my system message\"")
	}
}

func TestValidationErrors(t *testing.T) {
	type testStruct struct {
		name string
		age  int
	}

	value := &testStruct{name: "John", age: 20}
	errorValidation := validation.ValidateStruct(
		value,
		validation.Field(&value.name, validation.Required),
		validation.Field(&value.age, validation.Min(25)),
	)

	errMsg := NewError(0, errorValidation)
	errMsg.SetValidationMessage(errorValidation)

	assert.NotNil(t, errMsg.ValidationErrors["age"])
}

func TestErrorString(t *testing.T) {
	commonError := NewError(0, errors.New("this is internal error"))
	errMsg := commonError.Error()

	if assert.NotNil(t, errMsg) {
		assert.Contains(t, errMsg, "CommonError", "Trace")
	}
}
