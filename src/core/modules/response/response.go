package response

type BaseResponse[T any] struct {
	Data          T          `json:"data"`
	StatusCode    StatusCode `json:"status_code"`
	StatusMessage string     `json:"status_message"`
	TimeStamp     string     `json:"timeStamp"`
}

type StatusCode int16

const (
	StatusCodeGenericSuccess     StatusCode = 1000
	StatusCodeGenericClientError StatusCode = 2000
	StatusCodeGenericServerError StatusCode = 3000
	StatusCodeGenericHubError    StatusCode = 4000
)
