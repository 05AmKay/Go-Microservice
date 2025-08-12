package dto

import "time"

type ErrorResponseDto struct {
	ApiPath      string    `json:"apiPath"`
	ErrorMessage string    `json:"errorMessage"`
	ErrorCode    int       `json:"errorCode"`
	ErrorTime    time.Time `json:"errorTime"`
}

func NewErrorResponseDto(apiPath, errorMessage string, errorCode int) *ErrorResponseDto {
	return &ErrorResponseDto{
		ApiPath:      apiPath,
		ErrorMessage: errorMessage,
		ErrorCode:    errorCode,
		ErrorTime:    time.Now(),
	}
}
