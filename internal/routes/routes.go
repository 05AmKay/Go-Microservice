package routes

import (
	"example.com/api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	router.Use(middleware.GlobalErrorMiddleware())

	apiV1 := router.Group("/api/v1")

	// Register your routes here
	RegisterAccountRoutes(apiV1)

}
