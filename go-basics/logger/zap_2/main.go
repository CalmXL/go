package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func main() {

	/**
	  一个文件最多保存多大的日志体量  MaxSize
		最多备份多少个日志文件       MaxBackups
		老的日志文件最多保存多久     MaxAge
		要不要对备份的日志压缩       Compress
	*/
	infoLogger := &lumberjack.Logger{
		Filename:   "./logs/log-info.log",
		MaxSize:    1,  // mb
		MaxBackups: 3,  // 个文件
		MaxAge:     30, // days
		Compress:   true,
	}

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewConsoleEncoder(config)

	const (
		fileFlag = os.O_CREATE | os.O_WRONLY | os.O_APPEND
		prem     = 0666
	)

	fileMap := map[string]string{
		"DEBUG":  "./logs/log-debug.log",
		"INFO":   "./logs/log-info.log",
		"WARN":   "./logs/log-warn.log",
		"ERROR":  "./logs/log-error.log",
		"DPANIC": "./logs/log-dpanic.log",
		"PANIC":  "./logs/log-panic.log",
		"FATAL":  "./logs/log-fatal.log",
	}

	debugFile, _ := os.OpenFile(fileMap["DEBUG"], fileFlag, prem)
	// infoFile, _ := os.OpenFile(fileMap["INFO"], fileFlag, prem)
	warnFile, _ := os.OpenFile(fileMap["WARN"], fileFlag, prem)
	errorFile, _ := os.OpenFile(fileMap["ERROR"], fileFlag, prem)
	dpanicFile, _ := os.OpenFile(fileMap["DPANIC"], fileFlag, prem)
	panicFile, _ := os.OpenFile(fileMap["PANIC"], fileFlag, prem)
	fatalFile, _ := os.OpenFile(fileMap["FATAL"], fileFlag, prem)

	teeCore := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(debugFile), zapcore.DebugLevel),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(infoLogger), zapcore.InfoLevel),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(warnFile), zapcore.WarnLevel),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(errorFile), zapcore.ErrorLevel),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(dpanicFile), zapcore.DPanicLevel),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(panicFile), zapcore.PanicLevel),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(fatalFile), zapcore.FatalLevel),
	)

	logger := zap.New(teeCore, zap.AddCaller())

	defer logger.Sync()

	// logger.Debug("这是一个调试日志")
	// logger.Info("这是一个普通日志",
	// 	// zap.Namespace("medicine"),
	// 	zap.String("subject", "中医"),
	// 	zap.String("category", "中医药"),
	// )
	// logger.Warn("这是一个警告日志")
	// logger.Error("这是一个普通错误日志")
	// logger.DPanic("这是一个软错误日志")
	// logger.Panic("这是一个硬错误日志")
	// logger.Fatal("这是一个严重错误日志")

	for i := 0; i < 10000; i++ {
		logger.Debug("这是一个调试日志")
		logger.Info("这是一个普通日志",
			zap.Namespace("medicine"),
			zap.String("subject", "中医"),
			zap.String("category", "中医药"),
		)
		logger.Warn("这是一个警告日志")
		logger.Error("这是一个普通错误日志")
		// logger.DPanic("这是一个软错误日志")
		// logger.Panic("这是一个硬错误日志")
		// logger.Fatal("这是一个严重错误日志")
	}
}
