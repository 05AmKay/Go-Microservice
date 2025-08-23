package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"example.com/api/internal/config"
	"example.com/api/internal/models"
	"example.com/api/internal/routes"
	"example.com/api/internal/validation"
	"example.com/api/pkg/database"
	"example.com/api/pkg/logger"
	"github.com/gin-gonic/gin"
)

func init() {
	err := config.LoadAppConfig()
	if err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	if err := logger.InitializeLogger(); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
}

func main() {
	defer logger.Sync()
	Config := config.GetConfig()

	router := gin.Default()

	routes.RegisterRoutes(router)

	logger.Sugar.Infof("Setting up the database configuration...")
	err := database.InitializeDatabase()
	if err != nil {
		logger.Sugar.Fatalf("Failed to initialize database: %v", err)
		return
	}
	logger.Sugar.Infof("Database connection established successfully")

	db := database.GetDatabaseInstance().GetDB()
	if db == nil {
		logger.Sugar.Fatalf("Database connection is nil, initialization failed")
	}
	logger.Sugar.Infof("Database connection is ready to use. %T %v", db, db)

	dbModels := models.GetModels()
	if err := db.AutoMigrate(dbModels...); err != nil {
		logger.Sugar.Fatalf("failed to automigrate: %v", err)
	}
	logger.Sugar.Infof("Database table migration done successfully")

	logger.Sugar.Infof("Initializing the validator...")
	validation.InitValidator()
	logger.Sugar.Infof("Validator initialized successfully")

	// Create server
	srv := &http.Server{
		Addr:    ":" + Config.AppConfig.Port,
		Handler: router,
	}

	// Graceful shutdown setup (create a goroutine to start server and listen for shutdown signals later)
	go func() {
		logger.Sugar.Infof("Started server on port: %s", Config.AppConfig.Port)
		// Listen and serve
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Starting server failed: %s\n", err)
		}
	}()

	// Create a channerl to Wait for shutdown signal from goroutine
	quit := make(chan os.Signal, 1)

	// OS will notify the channel when an interrupt or terminate signal is received
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block the main goroutine until an interrupt or terminate signal is received
	<-quit
	logger.Sugar.Info("Shutdown Server ...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Sugar.Info("Server forced to shutdown:", err)
		log.Fatalf("Forced shutdown: %s\n", err)
	}
	logger.Sugar.Info("Server exiting")
}
