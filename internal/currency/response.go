package currency

import (
	"boletia/internal"
	"net/http"
)

type SuccessCurrency struct { //nolint: ineffassign
	data internal.Currencies
}

func NewSuccessCurrency(data internal.Currencies) SuccessCurrency {
	return SuccessCurrency{data: data}
}

func (h SuccessCurrency) Message(format string, data ...interface{}) string {
	return "success in obtaining the resource"
}

func (h SuccessCurrency) Code() int {
	return http.StatusOK
}

func (h SuccessCurrency) Error() string {
	return ""
}

func (h SuccessCurrency) Data() interface{} {
	return h.data
}
