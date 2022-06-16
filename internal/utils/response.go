package utils

type Response interface {
	Code() int
	Message(format string, data ...interface{}) string
	Error() string
	Data() interface{}
}

type HTTPResponse struct {
	Code    int    `json:"code,omitempty" example:"200"`
	Message string `json:"message,omitempty" example:"Success load the source"`
	Error   string `json:"error,omitempty" example:"error get some source"`
}
