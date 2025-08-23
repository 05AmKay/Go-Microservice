package exceptions

import (
	"net/http"

	errorfactory "example.com/api/internal/error"
	"github.com/gin-gonic/gin"
)

func ThrowError(c *gin.Context, appErr errorfactory.ApplicationError, cause ...error) {
	errorResponseDto := appErr.ToErrorResponseDto()
	c.AbortWithStatusJSON(errorResponseDto.ErrorCode, errorResponseDto)
}

func ThrowInternalServerError(c *gin.Context, cause ...error) {
	c.AbortWithStatusJSON(
		http.StatusInternalServerError,
		errorfactory.CreateInternalServerError(c.Request.URL.Path, cause...).ToErrorResponseDto())
}

func ThrowCustomError(c *gin.Context, statusCode int, errorMessage string, cause ...error) {
	c.AbortWithStatusJSON(
		statusCode,
		errorfactory.CreateCustomError(c.Request.URL.Path, statusCode, errorMessage, cause...).ToErrorResponseDto())
}
