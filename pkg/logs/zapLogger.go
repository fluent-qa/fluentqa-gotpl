package logs

import "go.uber.org/zap"

func NewDevLogger() *zap.Logger {
	logger, _ := zap.NewDevelopment()
	return logger
}

func NewProdLogger() *zap.Logger {
	logger, _ := zap.NewProduction()
	return logger
}
