package dto

import "time"

type ErrorResponseDto struct {
	ApiPath      string    `json:"api_path"`
	ErrorMessage any       `json:"error_message"`
	ErrorCode    int       `json:"error_code"`
	ErrorTime    time.Time `json:"time"`
}

func NewErrorResponseDto(apiPath string, errorMessage any, errorCode int) *ErrorResponseDto {
	return &ErrorResponseDto{
		ApiPath:      apiPath,
		ErrorMessage: errorMessage,
		ErrorCode:    errorCode,
		ErrorTime:    time.Now(),
	}
}
