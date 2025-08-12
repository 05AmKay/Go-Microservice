package errorfactory

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"example.com/api/internal/dto"
)

type AppError struct {
	ApiPath      string    `json:"apiPath"`
	ErrorMessage string    `json:"errorMessage"`
	ErrorCode    int       `json:"errorCode"`
	ErrorTime    time.Time `json:"errorTime"`
	Err          error     `json:"-"`
}

func NewAppError(apiPath, errorMessage string, errorCode int) *AppError {
	return &AppError{
		ApiPath:      apiPath,
		ErrorMessage: errorMessage,
		ErrorCode:    errorCode,
		ErrorTime:    time.Now(),
	}
}

func NewAppErrorWithCause(apiPath, errorMessage string, errorCode int, cause error) *AppError {
	ae := NewAppError(apiPath, errorMessage, errorCode)
	ae.Err = cause
	return ae
}

func IsAppError(err error) (*AppError, bool) {
	var appErr *AppError
	if ok := errors.As(err, &appErr); ok {
		return appErr, true
	}
	return nil, false
}

// Unwrap implements the errors.Unwrapper interface, crucial for errors.Is and errors.As.
func (e *AppError) Unwrap() error {
	return e.Err
}

// Error implements the error interface for AppError.
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("Path: %s, Code: %d, Message: %s, Wrapped Error: %s", e.ApiPath, e.ErrorCode, e.ErrorMessage, e.Err.Error())
	}
	return fmt.Sprintf("Path: %s, Code: %d, Message: %s", e.ApiPath, e.ErrorCode, e.ErrorMessage)
}

func (ae *AppError) ToErrorResponseDto() *dto.ErrorResponseDto {
	return &dto.ErrorResponseDto{
		ApiPath:      ae.ApiPath,
		ErrorMessage: ae.ErrorMessage,
		ErrorCode:    ae.ErrorCode,
		ErrorTime:    ae.ErrorTime,
	}
}

func NewResourceNotFoundError(resourceName, fieldName, fieldValue, apiPath string, cause ...error) *AppError {
	errorMessage := fmt.Sprintf("%s not found with %s: %v", resourceName, fieldName, fieldValue)
	if len(cause) > 0 {
		return NewAppErrorWithCause(apiPath, errorMessage, http.StatusNotFound, cause[0])
	}
	return NewAppError(apiPath, errorMessage, http.StatusNotFound)
}

func CustomerAlreadyExistsError(email, apiPath string, cause ...error) *AppError {
	errorMessage := fmt.Sprintf("Customer with email '%s' already exists", email)
	if len(cause) > 0 {
		return NewAppErrorWithCause(apiPath, errorMessage, http.StatusConflict, cause[0])
	}
	return NewAppError(apiPath, errorMessage, http.StatusConflict)
}
