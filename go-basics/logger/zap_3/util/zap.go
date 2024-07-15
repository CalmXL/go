package util

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
)

type LoggerConfig struct {
	Mode       string
	FilePath   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Level      zapcore.LevelEnabler
}

func NewLogger(config LoggerConfig) *zap.Logger {
	lumberLogger := &lumberjack.Logger{
		Filename:   config.FilePath,
		MaxSize:    config.MaxSize,
		MaxAge:     config.MaxAge,
		MaxBackups: config.MaxBackups,
		Compress:   true,
	}

	defer func(lumberLogger *lumberjack.Logger) {
		err := lumberLogger.Close()
		if err != nil {
			log.Fatal("Failed to close zap logger", zap.Error(err))
		}
	}(lumberLogger)

	var cfg zapcore.EncoderConfig

	switch config.Mode {
	case "DEV":
		cfg = zap.NewDevelopmentEncoderConfig()
	case "Prd":
		cfg = zap.NewProductionEncoderConfig()
	default:
		cfg = zap.NewDevelopmentEncoderConfig()
	}

	cfg.EncodeTime = zapcore.ISO8601TimeEncoder

	fileEncoder := zapcore.NewConsoleEncoder(cfg)

	core := zapcore.NewCore(fileEncoder, zapcore.AddSync(lumberLogger), config.Level)

	logger := zap.New(core, zap.AddCaller())

	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			log.Fatal("Failed to close zap logger", zap.Error(err))
			return
		}
	}(logger)

	zap.ReplaceGlobals(logger)

	return logger
}
