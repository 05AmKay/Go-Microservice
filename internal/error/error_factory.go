package errorfactory

import (
	"fmt"
	"net/http"
	"time"

	"example.com/api/internal/dto"
)

type ErrorType string

const (
	AppError                   ErrorType = "AppError"
	ValidationError            ErrorType = "ValidationError"
	ResourceNotFoundError      ErrorType = "ResourceNotFoundError"
	CustomerAlreadyExistsError ErrorType = "CustomerAlreadyExistsError"
)

type ApplicationError interface {
	error
	UnWrap() error
	Create(apiPath string, errorCode int, errorMessage any, cause ...error) ApplicationError
	ToErrorResponseDto() *dto.ErrorResponseDto
}

type BaseError struct {
	ApiPath   string    `json:"api_path"`
	ErrorCode int       `json:"error_code"`
	ErrorTime time.Time `json:"error_time"`
	Err       error     `json:"-"`
}
type AppErrorImpl struct {
	BaseError
	ErrorMessage string `json:"error_message"`
}

// Add UnWrap() method implementation for AppErrorImpl
func (ae *AppErrorImpl) UnWrap() error {
	return ae.Err
}

// Add Error() method implementation for AppErrorImpl
func (ae *AppErrorImpl) Error() string {
	if ae.Err != nil {
		return fmt.Sprintf("Path: %s, Code: %d, Message: %s, Wrapped Error: %s",
			ae.ApiPath, ae.ErrorCode, ae.ErrorMessage, ae.Err.Error())
	}
	return fmt.Sprintf("Path: %s, Code: %d, Message: %s",
		ae.ApiPath, ae.ErrorCode, ae.ErrorMessage)
}

func (ae *AppErrorImpl) Create(apiPath string, errorCode int, errorMessage any, cause ...error) ApplicationError {
	ae.ApiPath = apiPath
	ae.ErrorCode = errorCode
	ae.ErrorTime = time.Now()

	if len(cause) > 0 {
		ae.Err = cause[0]
	}

	if msg, ok := errorMessage.(string); ok {
		ae.ErrorMessage = msg
	}
	return ae
}

func GetErrorTypeFromFactory(errorType ErrorType) (ApplicationError, error) {
	switch errorType {
	case AppError, ResourceNotFoundError, CustomerAlreadyExistsError:
		return &AppErrorImpl{
			BaseError: BaseError{},
		}, nil
	case ValidationError:
		return &ValidationErrorImpl{
			BaseError:     BaseError{},
			ErrorMessages: []ValidationErrorDetail{},
		}, nil
	default:
		return nil, fmt.Errorf("unsupported database type: %s", errorType)
	}
}

func ThrowInternalServerError(apiPath string, cause ...error) ApplicationError {
	errObj, err := GetErrorTypeFromFactory(AppError)
	if err != nil {
		appError, _ := GetErrorTypeFromFactory(AppError)
		return appError.Create(
			apiPath, http.StatusInternalServerError, "Internal error factory failure", cause...)
	}
	return errObj.Create(apiPath, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), cause...)
}

func ThrowCustomError(apiPath string, errorCode int, errorMessage any, cause ...error) ApplicationError {
	errObj, err := GetErrorTypeFromFactory(AppError)
	if err != nil {
		return ThrowInternalServerError(apiPath, cause...)
	}
	return errObj.Create(apiPath, errorCode, errorMessage, cause...)
}

func ThrowCustomerAlreadyExistError(email, apiPath string, cause ...error) ApplicationError {
	errorMessage := fmt.Sprintf("Customer with email '%s' already exists", email)

	errObj, err := GetErrorTypeFromFactory(CustomerAlreadyExistsError)

	if err != nil {
		return ThrowInternalServerError(apiPath, cause...)
	}

	return errObj.Create(apiPath, http.StatusConflict, errorMessage, cause...)
}

// Add ToErrorResponseDto() method implementation for AppErrorImpl
func (ae *AppErrorImpl) ToErrorResponseDto() *dto.ErrorResponseDto {
	return &dto.ErrorResponseDto{
		ApiPath:      ae.ApiPath,
		ErrorMessage: ae.ErrorMessage,
		ErrorCode:    ae.ErrorCode,
		ErrorTime:    ae.ErrorTime,
	}
}
