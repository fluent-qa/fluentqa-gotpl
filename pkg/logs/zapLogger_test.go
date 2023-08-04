package logs

import (
	"go.uber.org/zap"
	"testing"
)

func TestNewDevLogger(t *testing.T) {
	tests := []struct {
		name string
		want *zap.Logger
	}{
		{name: "test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := NewDevLogger()
			logger.Info("test")
			logger.Error("test")
			logger.Warn("test")
			prodLogger := NewProdLogger()
			prodLogger.Info("test")
			prodLogger.Error("test")
			prodLogger.Warn("test")
		})
	}
}
