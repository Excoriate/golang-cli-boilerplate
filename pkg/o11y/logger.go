package o11y

import (
	"os"

	"go.uber.org/zap"
)

// NewLogger returns a new logger
func NewLogger() *zap.Logger {
	var logger *zap.Logger
	env := os.Getenv("ENV")
	if env == "prod" {
		logger, _ = zap.NewProduction()
		return logger
	}

	logger, _ = zap.NewDevelopment()
	return logger
}
