package main

import (
	"log"
	"logger/log_package/logger"
)

/*
*

	    日志: log => logger 日志记录工具
			常用的日志工具:
				logrus
				zerolog
				zap

			1. 操作
			2. 服务器程序的警告
			3. 调试
			4. 错误

		GO -> log 包

		日志的一般格式
			prefix                           DateTime               shortFilePath       info
			[INFO | WARN | ERROR | DEBUG]    2024-5-15 12:00:20     main.go(20行)       这是一个 log
*/
func main() {

	// 基本使用

	// 格式化打印 - format print
	// fmt.Println("这是一个格式化打印")
	// // 日志打印 - log print
	// log.Println("这是一个日志打印")

	// strFormat := "格式化"
	// strLog := "日志"
	//
	// fmt.Printf("这是一个%s打印\r\n", strFormat)
	// log.Printf("这是一个%s打印\r\n", strLog)

	// strFatal := "文件无法打开"
	// log.Fatalf("这是一个致命的错误: %s", strFatal)

	// strPainc := "程序出现意外"
	// // Painc 是会打印出错误栈的
	// log.Panicf("这时一个错误: %s", strPainc)

	// strLog := "日志"
	// log.SetPrefix("[INFO]")
	// log.Printf("这是一个%s打印\r\n", strLog)
	//
	// strFatal := "文件无法打开"
	// log.SetPrefix("[FATAL]")
	// log.Fatalf("这是一个致命的错误: %s", strFatal)

	// ----------------------------------------
	// Fatal 与 Painc 区别

	// log.Println("Print: 1")
	/**
	  func Fatalln(v ...any) {
		std.Output(2, fmt.Sprintln(v...))
		os.Exit(1) // 程序退出 成功 0 失败1
	  }

	  之前的同步代码执行， defer 不执行
	*/

	// defer func() {
	// 	log.Println("Print: defer1")
	// }()
	//
	// defer func() {
	// 	log.Println("Print: defer2")
	// }()
	//
	// defer func() {
	// 	log.Println("Print: defer3")
	// }()
	//
	// log.Fatalln("Fatal: 2")
	// log.Println("Println: 3")

	// log.Println("Print: 1")
	/**
	  func Panicln(v ...any) {
		s := fmt.Sprintln(v...)
		std.Output(2, s)
		panic(s)
	}

	panic
	  之前的同步代码执行，之后的同步代码被阻断
	  之前的defer 执行 之后的 defer 不执行
	*/

	// defer func() {
	// 	log.Println("Print: defer1")
	// }()
	//
	// defer func() {
	// 	log.Println("Print: defer2")
	// }()
	//
	// defer func() {
	// 	log.Println("Print: defer3")
	// }()
	//
	// log.Panicln("Painc: 2")
	// defer func() {
	// 	log.Println("Print: defer4")
	// }()
	// log.Println("Println: 3")

	// --------------------------------------------------
	/**
	  New 方法 返回一个 logger

		out io.Writer 输出对象 os.Stderr  os.Stdout  file
		prefix string, [info]
		flag int        额外的信息 123456
	*/
	/**
	  flag: 1 => [ERROR]2024/07/13 这是一个错误

	*/
	// log.Println(log.Ldate, log.Ltime, log.Lshortfile)
	// logger := log.New(os.Stderr, "[ERROR]", 2)
	// logger := log.New(os.Stdout, "[ERROR]", log.Ldate|log.Ltime|log.Lshortfile)

	// logger.Panicln("这是一个错误")

	err := logger.New()

	if err != nil {
		log.Fatalf("Failed to create log file: %s", err)
	}

	logger.Info("这是一个普通的日志：%s", "[Info]")
	logger.Warn("这是一个警告的日志%s", "[Warn]")
	logger.Error("这是一个错误的日志%s", "[Error]")
	logger.Debug("这是一个调试的日志%s", "[Debug]")
}
