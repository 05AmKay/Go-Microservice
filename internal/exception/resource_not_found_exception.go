package exceptions

import (
	"net/http"

	errorfactory "example.com/api/internal/error"
	"github.com/gin-gonic/gin"
)

func ResourceNotFoundException(c *gin.Context, resourceName, fieldName string, fieldValue any, casue ...error) {
	c.AbortWithStatusJSON(http.StatusNotFound, errorfactory.CreateResourceNotFoundError(
		c.Request.URL.Path,
		resourceName,
		fieldName,
		fieldValue,
		casue...,
	).ToErrorResponseDto())
}
