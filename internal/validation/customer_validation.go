package validation

import (
	"errors"
	"fmt"

	errorfactory "example.com/api/internal/error"
	exceptions "example.com/api/internal/exception"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func formatValidationErrorMessages(e validator.FieldError) errorfactory.ValidationErrorDetail {
	errorDetail := errorfactory.ValidationErrorDetail{
		Field: e.Field(),
	}

	switch e.Tag() {
	case "required":
		errorDetail.Error = fmt.Sprintf("Field '%s' is missing", e.Field())
	case "email":
		errorDetail.Error = fmt.Sprintf("Field '%s' must be a valid email", e.Field())
	case "gt":
		errorDetail.Error = fmt.Sprintf("Field '%s' must be greater than %s", e.Field(), e.Param())
	case "lt":
		errorDetail.Error = fmt.Sprintf("Field '%s' must be less than %s", e.Field(), e.Param())
	case "min":
		errorDetail.Error = fmt.Sprintf("Field '%s' must be at least %s characters long", e.Field(), e.Param())
	case "max":
		errorDetail.Error = fmt.Sprintf("Field '%s' must be at most %s characters long", e.Field(), e.Param())
	case "oneof":
		errorDetail.Error = fmt.Sprintf("Field '%s' must be one of the following values: %s", e.Field(), e.Param())
	case "gte":
		errorDetail.Error = fmt.Sprintf("Field '%s' must be greater than or equal to %s", e.Field(), e.Param())
	}

	return errorDetail
}

func ValidateCreateCustomerRequest(c *gin.Context, anyDto any) error {
	var validatorObj = GetValidator()

	// returns nil or ValidationErrors ( []FieldError )
	err := validatorObj.Struct(anyDto)
	if err != nil {
		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			fmt.Println(err)
			return fmt.Errorf("validation failed, invalid value for validation")
		}

		var validateErrs validator.ValidationErrors
		if errors.As(err, &validateErrs) {
			errorMessages := make([]errorfactory.ValidationErrorDetail, 0)
			for _, e := range validateErrs {
				errorDetail := formatValidationErrorMessages(e)
				errorMessages = append(errorMessages, errorDetail)

				// fmt.Println("Namespace: ", e.Namespace())
				// fmt.Println("Field: ", e.Field())
				// fmt.Println("StructNamespace: ", e.StructNamespace())
				// fmt.Println("StructField: ", e.StructField())
				// fmt.Println("Tag: ", e.Tag())
				// fmt.Println("ActualTag: ", e.ActualTag())
				// fmt.Println("Kind: ", e.Kind())
				// fmt.Println("Typee: ", e.Type())
				// fmt.Println("Value: ", e.Value())
				// fmt.Println("Param: ", e.Param())
				// fmt.Println()
			}
			fmt.Println("Validation errors:", errorMessages)
			exceptions.ThrowValidationException(c, errorMessages)
			return fmt.Errorf("validation failed")
		}

	}
	return nil
}
