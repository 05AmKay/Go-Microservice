package exceptions

import (
	"net/http"

	errorfactory "example.com/api/internal/error"
	"github.com/gin-gonic/gin"
)

func ThrowValidationException(c *gin.Context, errorMessages any, cause ...error) {
	c.AbortWithStatusJSON(http.StatusBadRequest, errorfactory.CreateValidationError(
		c.Request.URL.Path,
		errorMessages,
		cause...,
	).ToErrorResponseDto())
}
