package logger

import "testing"

func TestInfo(t *testing.T) {
	Info("test")
}

func TestNewZapLogger(t *testing.T) {
	logger := NewZapLogger()
	logger.Info("info", "asdfadfsadfdsafds")
	logger.Error("error", "asdfadfsadfdsafds")
	logger.Debug("debug", "asdfadfsadfdsafds")
}
