package main

import (
	"fmt"
	"sync"
	"time"
)

// goroutine
/**
当 main 函数内部执行完所有函数后，直接退出
不会等待异步任务完成
*/
func main() {

	var waitGroup sync.WaitGroup
	// // 创建并执行一个 Goroutine 任务
	// // 一个异步任务
	// // test0 开始执行，则 test1 执行
	// // test0 是否执行完毕，与 test1 执行无关
	// go test0()
	// // 0 => 是在同步任务全部执行完毕后，再输出吗？
	// // 非阻塞状态
	// /*
	//  这个 Goroutine 什么时候完成是未知的
	// */
	//
	// // 同步任务执行的过程
	// // 一个函数执行完成之后，在执行下一个函数(串行方式)
	// test1()
	// // 阻塞
	// test2()
	// test3()
	//
	// for i := 0; i < 100; i++ {
	// 	fmt.Println(i)
	// }

	// ***********************************************

	for i := 0; i < 100; i++ {
		/**
		Goroutine 的 Processor 调度机制决定了每个 Goroutine 的执行时机
		*/
		// go func() {
		// 	fmt.Println(i)
		// }()

		// 临时变量机制
		// t := i
		// go func() {
		// 	fmt.Println(t)
		// }()

		// 同步等待组
		waitGroup.Add(1)

		// 值复制机制
		// 参数传递机制
		go func(i int) {
			// waitGroup.Add(-1)
			// deadLock 死锁
			defer waitGroup.Done()
			fmt.Println(i)
		}(i)
	}

	time.Sleep(time.Second)
}

func test0() {
	fmt.Println("test0 => 0")
}

func test1() {
	fmt.Println(1)

}

func test2() {
	fmt.Println(2)

}

func test3() {
	fmt.Println(3)

}
