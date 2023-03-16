package logger

import (
	"encoding/json"
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger instantiates a new zap.SugaredLogger.
func NewLogger() *zap.SugaredLogger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "time"
	config.EncoderConfig.TimeKey = "level"
	config.EncoderConfig.MessageKey = "msg"
	config.Encoding = "console"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.DisableCaller = true
	config.DisableStacktrace = true
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncoderConfig.ConsoleSeparator = "  "
	logger, err := config.Build()
	if err != nil {
		log.Fatalf("Creating logger. %v", err)
	}
	return logger.Sugar()
}

// LogSync synchronizes a zap.SugaredLogger.
func LogSync(log *zap.SugaredLogger) {
	_ = log.Sync()
}

// LogLevel returns the definition of levels in logs.
func LogLevel(logLevel string) zapcore.Level {
	var lvl zapcore.Level
	err := json.Unmarshal([]byte("\""+logLevel+"\""), &lvl)
	if err != nil {
		return zapcore.InfoLevel
	}
	return lvl
}
