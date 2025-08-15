package errorfactory

import (
	"fmt"
	"net/http"
	"time"

	"example.com/api/internal/dto"
)

type ValidationErrorDetail struct {
	Field string `json:"field"`
	Error string `json:"error"`
}
type ValidationErrorImpl struct {
	BaseError
	ErrorMessages []ValidationErrorDetail `json:"error_message"`
}

// Add UnWrap() method implementation for ValidationErrorImpl
func (ve *ValidationErrorImpl) UnWrap() error {
	return ve.Err
}

// Add Error() method implementation for ValidationErrorImpl
func (ve *ValidationErrorImpl) Error() string {
	if ve.Err != nil {
		return fmt.Sprintf("Path: %s, Code: %d, Validations: %v, Wrapped Error: %s",
			ve.ApiPath, ve.ErrorCode, ve.ErrorMessages, ve.Err.Error())
	}
	return fmt.Sprintf("Path: %s, Code: %d, Validations: %v",
		ve.ApiPath, ve.ErrorCode, ve.ErrorMessages)
}

func (ve *ValidationErrorImpl) Create(apiPath string, errorCode int, errorMessage any, cause ...error) ApplicationError {
	ve.ApiPath = apiPath
	ve.ErrorCode = errorCode
	ve.ErrorTime = time.Now()

	if len(cause) > 0 {
		ve.Err = cause[0]
	}

	if msgs, ok := errorMessage.([]ValidationErrorDetail); ok {
		ve.ErrorMessages = msgs
	}
	return ve
}

func ThrowValidationError(apiPath string, errorMessage any, errorCode int, cause ...error) ApplicationError {
	errObj, err := GetErrorTypeFromFactory(ValidationError)
	if err != nil {
		return ThrowInternalServerError(apiPath, cause...)
	}

	return errObj.Create(apiPath, http.StatusBadRequest, errorMessage, cause...)
}

// Add ToErrorResponseDto() method implementation for ValidationErrorImpl
func (ve *ValidationErrorImpl) ToErrorResponseDto() *dto.ErrorResponseDto {
	return &dto.ErrorResponseDto{
		ApiPath:      ve.ApiPath,
		ErrorMessage: ve.ErrorMessages,
		ErrorCode:    ve.ErrorCode,
		ErrorTime:    ve.ErrorTime,
	}
}
