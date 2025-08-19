package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Generate a unique request ID and set it in the context
		requestID := c.Request.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = "req-" + uuid.New().String()
		}

		c.Set("requestID", requestID)

		c.Next()
	}
}
