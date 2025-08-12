package main

import (
	"fmt"

	"example.com/api/pkg/database"
)

func main() {
	fmt.Println("Setting up the database configuration...")

	configBuilder := database.NewDbConfigBuilder()
	config, err := configBuilder.SetHost("localhost").
		SetPort(5432).
		SetCredentials("postgres", "password").
		SetDatabase("postgres").
		SetSSL(false).
		Build()

	if err != nil {
		fmt.Println("Error building database config:", err)
		return
	}
	fmt.Printf("Database configuration built successfully: %v\n", config)

	dbConn, err := database.GetDatabaseConnectionFromFactory(database.Postgres, config)
	if err != nil {
		fmt.Println("Error getting database connection from factory:", err)
		return
	}
	fmt.Printf("Database connection established successfully %v\n", dbConn)

}
