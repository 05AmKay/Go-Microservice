package dto

type ResponseDto struct {
	StatusCode    int    `json:"status"`
	StatusMessage string `json:"message"`
	Data          any    `json:"data,omitempty"`
}

func NewResponseDto(statusCode int, statusMessage string, dto any) *ResponseDto {
	if dto != nil {
		return &ResponseDto{
			StatusCode:    statusCode,
			StatusMessage: statusMessage,
			Data:          dto,
		}
	} else {
		return &ResponseDto{
			StatusCode:    statusCode,
			StatusMessage: statusMessage,
			Data:          nil,
		}
	}
}
