package test

import "go.uber.org/zap"

func Test() {
	zap.L().Info("这一个 test 内部的普通日志")
	zap.S().Warn("这是一个警告日志")
}
