package utils

import (
	"fmt"
	"net/http"
)

type InternalResponseError struct {
	err  error
	data interface{}
}

func NewInternalErrResponse(err error, data interface{}) *InternalResponseError {
	return &InternalResponseError{err: err, data: data}
}

func (i InternalResponseError) Code() int {
	return http.StatusInternalServerError
}

func (i InternalResponseError) Message(format string, data ...interface{}) string {
	if format != "" {
		return fmt.Sprintf(format, data)
	}
	if len(data) > 0 {
		return fmt.Sprintf("Error getting the source %s", data)
	}
	return fmt.Sprintf("The server not responsed correctly")
}

func (i InternalResponseError) Error() string {
	return i.err.Error()
}

func (i InternalResponseError) Data() interface{} {
	return i.data
}
