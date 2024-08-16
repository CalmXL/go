package util

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"sale_system/config"
)

type LoggerConfig struct {
	Mode       string
	FilePath   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Level      zapcore.LevelEnabler
}

func NewLogger(cfg *config.LogConfig) *zap.Logger {
	env := config.GetEnv(config.FileConfig["ENV"])

	lumberLogger := &lumberjack.Logger{
		Filename:   cfg.FilePath,
		MaxSize:    cfg.MaxSize,
		MaxAge:     cfg.MaxAge,
		MaxBackups: cfg.MaxBackups,
		Compress:   true,
	}

	defer func(lumberLogger *lumberjack.Logger) {
		err := lumberLogger.Close()
		if err != nil {
			log.Fatal("Failed to close zap logger", zap.Error(err))
		}
	}(lumberLogger)

	_cfg := zap.NewDevelopmentEncoderConfig()

	if env {
		_cfg = zap.NewDevelopmentEncoderConfig()
	} else {
		_cfg = zap.NewProductionEncoderConfig()
	}

	_cfg.EncodeTime = zapcore.ISO8601TimeEncoder

	fileEncoder := zapcore.NewConsoleEncoder(_cfg)

	core := zapcore.NewCore(fileEncoder, zapcore.AddSync(lumberLogger), config.ZapLevel)

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
