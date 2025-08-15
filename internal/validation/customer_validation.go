package validation

import (
	"errors"
	"fmt"
	"net/http"

	"example.com/api/internal/dto"
	errorfactory "example.com/api/internal/error"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidateCreateCustomerRequest(c *gin.Context, customerDto dto.CustomerDto) {
	var validatorObj = GetValidator()

	// returns nil or ValidationErrors ( []FieldError )
	err := validatorObj.Struct(customerDto)
	if err != nil {
		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			fmt.Println(err)
			return
		}

		var validateErrs validator.ValidationErrors
		if errors.As(err, &validateErrs) {
			errorMessages := make([]errorfactory.ValidationErrorDetail, 0)
			for _, e := range validateErrs {
				errorDetail := errorfactory.ValidationErrorDetail{
					Field: e.Field(),
				}
				switch e.Tag() {
				case "required":
					errorDetail.Error =
						fmt.Sprintf("Field '%s' is missing", e.Field())
				case "email":
					errorDetail.Error =
						fmt.Sprintf("Field '%s' must be a valid email", e.Field())
				case "gte":
					errorDetail.Error =
						fmt.Sprintf("Field '%s' must be greater than or equal to %s", e.Field(), e.Param())
				}

				errorMessages = append(errorMessages, errorDetail)

				fmt.Println("Namespace: ", e.Namespace())
				fmt.Println("Field: ", e.Field())
				fmt.Println("StructNamespace: ", e.StructNamespace())
				fmt.Println("StructField: ", e.StructField())
				fmt.Println("Tag: ", e.Tag())
				fmt.Println("ActualTag: ", e.ActualTag())
				fmt.Println("Kind: ", e.Kind())
				fmt.Println("Typee: ", e.Type())
				fmt.Println("Value: ", e.Value())
				fmt.Println("Param: ", e.Param())
				fmt.Println()
			}
			fmt.Println("Validation errors:", errorMessages)
			c.Error(errorfactory.ThrowValidationError(
				c.Request.URL.Path,
				errorMessages,
				http.StatusBadRequest,
			))
			return
		}

	}

}
