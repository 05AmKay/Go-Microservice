package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger
var Sugar *zap.SugaredLogger

func InitializeLogger() error {

	var err error

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}
	switch env {
	case "production":
		logger, err = zap.NewProduction()
	case "development":
		logger, err = zap.NewDevelopment()
	default:
		// Custom configuration
		config := zap.NewProductionConfig()
		config.EncoderConfig.TimeKey = "timestamp"
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		logger, err = config.Build()

	}

	if err != nil {
		return err
	}

	Sugar = logger.Sugar()

	return nil
}

func Sync() error {
	if logger != nil {
		return logger.Sync() // Flushes any buffered log entries
	}
	return nil
}
