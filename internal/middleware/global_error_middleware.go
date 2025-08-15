package middleware

import (
	"log"
	"net/http"

	"example.com/api/internal/dto"
	errorfactory "example.com/api/internal/error"
	"github.com/gin-gonic/gin"
)

func GlobalErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Process the request and subsequent handlers in the chain

		// Check if any errors were added to the context by handlers.
		if len(c.Errors) > 0 {
			lastError := c.Errors.Last().Err

			log.Printf("An  error occured: %v", lastError)
			var errorResponseDto *dto.ErrorResponseDto

			if appErr, ok := lastError.(errorfactory.ApplicationError); ok {
				errorResponseDto = appErr.ToErrorResponseDto()
			} else {
				appErr = errorfactory.ThrowInternalServerError(
					c.Request.URL.Path,
				)
				errorResponseDto = appErr.ToErrorResponseDto()
			}

			// Do not overwrite an existing status code if set by AbortWithError
			// Or explicitly set status code and JSON response
			if c.Writer.Status() == http.StatusOK { // Only set status if not already set
				c.AbortWithStatusJSON(errorResponseDto.ErrorCode, errorResponseDto) //
			} else {
				// If a status was already set (e.g., by AbortWithError), update the JSON payload only
				c.AbortWithStatusJSON(c.Writer.Status(), errorResponseDto) //
			}
		}

	}
}
