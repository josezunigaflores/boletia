package utils

import (
	"fmt"
	"net/http"
)

type BadRequestError struct {
	err  error
	data interface{}
}

func NewBadRequest(err error, data ...interface{}) *BadRequestError {
	return &BadRequestError{err: err, data: data}
}

func (b BadRequestError) Code() int {
	return http.StatusBadRequest
}

func (b BadRequestError) Message(format string, data ...interface{}) string {
	if format != "" {
		return fmt.Sprintf(format, data...)
	}
	if len(data) > 0 {
		return fmt.Sprintf("Bad request by %s", data)
	}

	return "Bad request"
}

func (b BadRequestError) Error() string {
	return b.err.Error()
}

func (b BadRequestError) Data() interface{} {
	return b.data
}
