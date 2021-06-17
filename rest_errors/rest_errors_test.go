package rest_errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("this is the message", errors.New("database error"))

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "this is the message", err.Message)
	assert.EqualValues(t, "internal_server_error", err.Error)
	assert.NotNil(t, err.Causes)
	errByte, _ := json.Marshal(err)
	fmt.Println(string(errByte))
}

func TestNewBadRequestError(t *testing.T) {
	// Todo: test
}

func TestNewNotFoundError(t *testing.T) {
	// Todo: test
}

func TestNewError(t *testing.T) {
	// Todo: test
}
