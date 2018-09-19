package utils

import (
	"fmt"
	"go.uber.org/zap"
)

var (
	Logger = New()
)

func New() *zap.Logger {
	fmt.Println("Hello")
	
	logger, _ := zap.NewProduction()
	logger.Sync() // flushes buffer, if any
	
	return logger
}

