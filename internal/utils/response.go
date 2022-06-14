package utils

type Response interface {
	Code() int
	Message(format string, data ...interface{}) string
	Error() string
	Data() interface{}
}

type HTTPResponse struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data"`
}
