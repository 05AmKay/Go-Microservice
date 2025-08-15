package dto

type ResponseDto struct {
	StatusCode    int    `json:"status"`
	StatusMessage string `json:"message"`
}

func NewResponseDto(statusCode int, statusMessage string) *ResponseDto {
	return &ResponseDto{
		StatusCode:    statusCode,
		StatusMessage: statusMessage,
	}
}
