package main

import (
	"fmt"
	"log"

	"example.com/api/internal/routes"
	"example.com/api/internal/validation"
	"example.com/api/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.RegisterRoutes(router)

	fmt.Println("Setting up the database configuration...")
	err := database.InitializeDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
		return
	}
	fmt.Printf("Database connection established successfully")
	db := database.GetDatabaseInstance().GetDB()
	if db == nil {
		fmt.Println("Database connection is nil, initialization failed.")
	}
	fmt.Printf("Database connection is ready to use. %T %v", db, db)

	fmt.Println("Initializing the validator...")
	validation.InitValidator()
	fmt.Println("Validator initialized successfully.")

	log.Println("Server starting on :8080")
	log.Fatal(router.Run(":8080"))

}
