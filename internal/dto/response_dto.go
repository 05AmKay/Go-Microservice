package dto

type ResponseDto struct {
	StatusCode    string `json:"statusCode"`
	StatusMessage string `json:"statusMessage"`
}

func NewResponseDto(statusCode, statusMessage string) *ResponseDto {
	return &ResponseDto{
		StatusCode:    statusCode,
		StatusMessage: statusMessage,
	}
}
