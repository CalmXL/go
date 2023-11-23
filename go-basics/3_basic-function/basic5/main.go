package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
  defer、panic、recover

	错误机制
		1. 语法错误
		2. 运行错误
		3. 系统操作错误

	错误抛出和捕获机制

		抛出异常：当前这个错误（异常）我需要有明确的说明
			throw Error("文件读取失败", err)

		捕获异常: 一旦出现系统操作的错误，那么整个程序就会停止工作
		捕获一些我们没办法控制是否会出错的错误

		如果不捕获错误：未捕获的错误（异常）

		readFile
			=> 未捕获这个行为的异常
			--- 中断 ---
			res := a + b


		try {
      // 读取文件
			=> 读取成功
					=> 一行一行的输出
		} catch(err) {
			=> 读取失败
				进行读取失败的逻辑
		} finally {
			=> 不管读取文件是否失败
				都会执行这里的逻辑
		}

	go 中：
		抛出异常：painc
		捕获异常：recover`
    最终逻辑：defer

	error:
		Go 语言中，不认为所有的错误都需要用捕获的方式来获取错误信息
		每个需要可能抛出错误的方法都应该返回一个 error

*/

/**
  defer 语句
	defer 常常跟的是函数执行
	延迟函数执行的时机
	前面跟 defer 语句都会在函数 return 之后或函数最后进行执行
	多个 refer 的函数倒序处理

	主要是处理函数最后必要完成的任务
	读取文件、连接数据库、锁相关的操作
*/

// func test() (int, error) {
// 	return 1, err
// }

func main() {
	// res, error := test()

	// defer fmt.Println("defer1")
	// defer fmt.Println("defer2")
	// defer fmt.Println("defer3")
	//
	// fmt.Println("main1")
	// fmt.Println("main2")
	// fmt.Println("main3")

	// test()

	readFile()
	fmt.Println("Continuous program")
}

func readFile() {
	defer catchPainc()
	file, err := os.Open("./_basic-function/basic5/desc.txt")
	defer func(file *os.File) {
		fmt.Println(92)
		if err := file.Close(); err != nil {
			panic("关闭文件失败")
		}
	}(file)

	fmt.Println(97, err)
	if err != nil {
		fmt.Println(99, err)
		panic("打开文件失败")
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}

func catchPainc() {
	if err := recover(); err != nil {
		fmt.Println("catch panic", err)
	}
}

// func test() int {
// 	defer fmt.Println("defer1")
// 	defer fmt.Println("defer2")
// 	defer fmt.Println("defer3")
//
// 	fmt.Println("test1")
// 	fmt.Println("test2")
// 	fmt.Println("test3")
//
// 	// return 的目的
// 	/**
// 	  1. 返回一个值
// 		2. 终止函数程序执行
// 	*/
// 	return 1
// }
