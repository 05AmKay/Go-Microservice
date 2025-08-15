package errorfactory

import (
	"fmt"
	"net/http"
)

func ThrowResourceNotFoundError(resourceName, fieldName, fieldValue, apiPath string, cause ...error) ApplicationError {
	errorMessage := fmt.Sprintf("%s not found with %s: %v", resourceName, fieldName, fieldValue)

	errObj, err := GetErrorTypeFromFactory(ResourceNotFoundError)

	if err != nil {
		return ThrowInternalServerError(apiPath, cause...)
	}

	return errObj.Create(apiPath, http.StatusNotFound, errorMessage, cause...)
}
