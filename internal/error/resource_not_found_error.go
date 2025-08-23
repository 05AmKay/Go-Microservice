package errorfactory

import (
	"fmt"
	"net/http"
)

func CreateResourceNotFoundError(apiPath string, resourceName, fieldName string, fieldValue any, cause ...error) ApplicationError {
	errorMessage := fmt.Sprintf("%s not found with %s: %v", resourceName, fieldName, fieldValue)

	errObj, err := GetErrorTypeFromFactory(ResourceNotFoundError)

	if err != nil {
		return CreateInternalServerError(apiPath, cause...)
	}

	return errObj.Create(apiPath, http.StatusNotFound, errorMessage, cause...)
}
