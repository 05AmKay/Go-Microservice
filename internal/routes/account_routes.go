package routes

import (
	"example.com/api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterAccountRoutes(apiV1 *gin.RouterGroup) {
	account := apiV1.Group("/account")
	{
		account.POST("/create", handlers.CreateAccount)
	}
}
