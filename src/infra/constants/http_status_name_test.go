package constants

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisteredHttpStatusText(t *testing.T) {
	httpStatusText := GetHttpStatusText(400)

	assert.Equal(t, "BAD_REQUEST", httpStatusText, "they should be equal BAD_REQUEST")
}

func TestUnregisteredHttpStatusText(t *testing.T) {
	httpStatusText := GetHttpStatusText(4000)

	assert.Equal(t, "INTERNAL_SERVER_ERROR", httpStatusText, "they should be equal INTERNAL_SERVER_ERROR")
}
