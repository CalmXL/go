package main

import (
	"fmt"
	"go.uber.org/zap/zapcore"
	"logger/zap_3/test"
	"logger/zap_3/util"
)

/*
*

	  Zap 两个大的 logger 类型
		1. 是由 New 创建 logger 对象，利用 logger 对象来输出日志
			logger.Info()
		2. 用 zap.L() 和 zap.S()  创建标准 logger 和 SugaredLogger(全局logger)
			logger
			zap.ReplaceGlobals(logger)

			zap.L().Info()
			zap.S().Info()
*/
func main() {
	// logger, _ := zap.NewProduction()
	//
	// zap.ReplaceGlobals(logger)

	logger := util.NewLogger(util.LoggerConfig{
		Mode:       "Prd",
		FilePath:   "./logs/log-ingo.log",
		MaxSize:    1,
		MaxBackups: 3,
		MaxAge:     28,
		Level:      zapcore.InfoLevel,
	})

	// logger.Info("这是一个普通日志")
	// logger.Warn("这是一个警告日志")
	//
	// zap.L().Info("这是一个普通日志")
	// zap.S().Warn("这是一个警告日志")
	fmt.Println(logger)

	for i := 0; i < 10000; i++ {
		test.Test()
	}
}
