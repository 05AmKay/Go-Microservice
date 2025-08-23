package routes

import (
	"example.com/api/internal/handlers"
	"example.com/api/internal/service/impl"
	"github.com/gin-gonic/gin"
)

func RegisterAccountRoutes(apiV1 *gin.RouterGroup) {
	// Initialize service
	accountService := &impl.AccountServiceImpl{}

	accountHandler := handlers.NewAccountHandler(accountService)

	account := apiV1.Group("/account")
	{
		account.POST("/create", accountHandler.CreateAccount)
		account.GET("/fetch", accountHandler.FetchAccountDetails)
		account.PUT("/update", accountHandler.UpdateAccountDetails)
		account.DELETE("/delete", accountHandler.DeleteAccount)
	}
}
