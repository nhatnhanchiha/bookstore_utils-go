package rest_errors

import (
	"errors"
	"fmt"
	"net/http"
)

type restErr struct {
	ErrMessage string        `json:"message"`
	ErrStatus  int           `json:"status"`
	ErrError   string        `json:"error"`
	ErrCauses  []interface{} `json:"causes"`
}

func (e restErr) Message() string {
	return e.ErrMessage
}

func (e restErr) Status() int {
	return e.ErrStatus
}

func (e restErr) Causes() []interface{} {
	return e.ErrCauses
}

type RestErr interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

func (e restErr) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes [ %v ]", e.ErrMessage, e.ErrStatus, e.ErrError, e.ErrCauses)
}

func NewError(msg string) error {
	return errors.New(msg)
}

func NewBadRequestError(message string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "bad_request",
	}
}

func NewNotFoundError(message string) RestErr {
	return &restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "not_found",
	}
}

func NewInternalServerError(message string, err error) RestErr {
	result := restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "internal_server_error",
	}

	if err != nil {
		result.ErrCauses = append(result.ErrCauses, err.Error())
	}

	return result
}

func NewUnauthorizedError(message string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusUnauthorized,
		ErrError:   "unauthorized",
	}
}

func NewRestError(message string, status int, error string, causes []interface{}) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  status,
		ErrError:   error,
		ErrCauses:  causes,
	}
}
