package main

import (
	"fmt"
	"log"

	errorfactory "example.com/api/internal/error"
	"example.com/api/internal/middleware"
	"example.com/api/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.GlobalErrorMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.Error(errorfactory.NewResourceNotFoundError("User", "username", "test", c.Request.URL.Path))
	})

	fmt.Println("Setting up the database configuration...")
	database.InitializeDatabase()
	fmt.Printf("Database connection established successfully")
	db := database.GetDatabaseInstance().GetDB()
	if db == nil {
		fmt.Println("Database connection is nil, initialization failed.")
	}
	fmt.Printf("Database connection is ready to use. %T %v", db, db)

	log.Println("Server starting on :8080")
	log.Fatal(r.Run(":8080"))

}
