package util

import (
	"log"

	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	logger = zapLogger
}

func LoggerInstance() *zap.Logger {
	return logger
}
