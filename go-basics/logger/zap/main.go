package main

import (
	"encoding/json"
	"go.uber.org/zap"
)

/*
*

	  Zap: Uber 开源的日志记录包

		高性能
		日志等级的封装很科学
		支持自定义的日志格式
		支持公共字段的定义
		细节化调试字段
		调用栈信息输出
	    日志命名空间
		NameSpace {
			a: {}
			b: {}
		}

		支持钩子操作 Hook
*/
func main() {
	/**
	  1. Zap 分为两种 logger
		1. 标准 logger
			日志的相关字段是强类型的
	        性能非常高
			更加的安全
		2. 日志糖 SugaredLogger
			配置字段弱类型是分散的
			性能略低于 logger
			安全性略低
			支持Printf
		3. Zap  创建 Logger 对象的方式三种：
			开发模式: NewDevelopment
				会输出一个标准的 logger 模式
			生产模式: NewProduction
				会输出一个标准的 JSON 格式
			示例模式: NewExample
				没有时间文件位置
				只有等级和信息
				输出一个标准的 JSON 模式的字符串
	   2. Zap 的 日志等级
			1. Debug => 调试日志 （Production 模式下不包含）
			2. Info => 普通日志
			3. Warn => 警告日志
			4. Error => 普通错误日志
			5. DPanic => 软 Panic（错误）日志
			6. Painc =>  硬Panic (错误) 日志
			7. Fatal => 严重错误日志
	*/

	// logger, err := zap.NewDevelopment()
	// logger, err := zap.NewProduction()
	// logger := zap.NewExample()

	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	//
	// defer logger.Sync() // flushes buffer
	//
	// logger.Debug("这是一个调试日志")
	// logger.Info("这是一个普通的日志")
	// logger.Warn("这是一个警告的日志")

	// logger.DPanic("这是一个Dpanic 错误日志")
	// logger.Panic("这是一个 Panic 错误日志")
	// logger.Error("这是一个普通错误日志")
	// logger.Fatal("这是一个严重错误日志")

	/**
	    Error: 错误事件，只是一个普通的错误，程序是可以继续运行的
		DPanic: 软错误，输出错误栈信息，程序不会终止
		Panic: 严重错误，输出错误栈信息，程序会终止
		Fatal: 严重错误，不会输出错误栈信息，程序会终止
	*/

	// -----------------------------------------------------
	// logger, err := zap.NewProduction()
	//
	// if err != nil {
	// 	log.Fatal("logger init failed", zap.Error(err))
	// 	return
	// }
	//
	// defer logger.Sync()

	// su := logger.Sugar()

	// su.Debugf("这是一个调试日志 %s", "[Debug]")
	// su.Fatalw("failed to connect to database: ",
	// 	// 松散性 key value 传参方式且弱类型
	// 	"IP", "192.168.101.14",
	// 	"Port", 3006,
	// )

	// logger.Fatal("Failed to connect to database: ",
	// 	zap.String("IP", "192.168.101.14"),
	// 	zap.Int("Port", 3006),
	// )

	// -------------------------------------------
	// zap.Hooks
	// 当一次日志记录完毕后，调用 zap.Hooks 里的回调函数
	// logger, _ := zap.NewProduction(zap.Hooks(func(entry zapcore.Entry) error {
	//
	// 	fmt.Println("entry: ", entry)
	// 	return nil
	// }, func(entry zapcore.Entry) error {
	// 	fmt.Println("entry2: ", entry)
	// 	return nil
	// }))
	//
	// logger.Info("这是一个普通的日志")
	// logger.Error("这是一个错误日志")

	// ---------------------------------------------

	rawStream := []byte(`
		{
			"level": "debug",
			"encoding": "json",
			"outputPaths": ["stdout", "./logs/log-info.log"],
			"errorOutputPaths": ["stderr", "./logs/log-error.log"],
			"encoderConfig": {
				"messageKey": "msg",
				"levelKey": "level",
				"levelEncoder": "capital"
			}
		}
	`)

	var config zap.Config
	if err := json.Unmarshal(rawStream, &config); err != nil {
		panic(err)
	}

	logger := zap.Must(config.Build())

	defer logger.Sync()
	//
	// logger.Info("这是一个普通日志")
	// logger.Error("这是一个错误日志")
	// logger.Panic("这时一个错误日志")
	// logger.Fatal("这是一个严重的错误日志")
	//
	su := logger.Sugar()

	su.Warnf("注意有新的API可以替换：%s", "abc -> bcd")
	su.Fatalf("failed to fetch url: %s", "https://www.baidu.com")
}
