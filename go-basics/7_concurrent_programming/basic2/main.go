package main

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
)

/**
  原子操作：一个程序开始运行直到运行完毕后，再进行下一个任务程序的执行的操作。
	在 Go 语言中，GoRoutine 们的运行针对同一个变量的操作是竞争态的关系

	竞争态 => 数据竞争 Date Race / Race Status
	竞争态关系会使变量的结果变得不可预料
	所以在某个 Goroutine 操作一个变量的时候，一定需要一种锁定机制
	让系统禁止多个 Goroutine 操作同一个变量
	如果不禁止， Goroutine 的运行会使程序执行变得不安全

	Go 语言中: Mutex => 互斥锁
*/
var a int32 = 0
var wg sync.WaitGroup

/**
  1. 不要复制 Mutex 锁对象
	2. 需要解决竞争态的锁分配的多个 Goroutine 使用同一个 Mutex 对象
		 因为 Mutex 会产生 Goroutine 运行得阻塞
     如果没有竞争态或不需要解决竞争关系的多个 Goroutine 不要使用同一把锁
*/
var lk sync.Mutex

func main() {
	wg.Add(2)
	go plus()
	go minus()

	wg.Wait()
	fmt.Println(a)
}

func plus() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		// lk.Lock()
		// a += 1
		// lk.Unlock()

		// 原子操作的 Add 方法操作的是指针对应的空间，直接修改值
		atomic.AddInt32(&a, 1)
		fmt.Println("plus: " + strconv.Itoa(int(a)))
	}
}

func minus() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		// lk.Lock()
		// a -= 1
		// lk.Unlock()

		// 原子操作的 Add 方法操作的是指针对应的空间，直接修改值
		atomic.AddInt32(&a, -1)
		fmt.Println("minus: " + strconv.Itoa(int(a)))
	}
}

/**
  Mutex =>
		state: 互斥锁状态
				Locked
				Woken: 唤醒的
				Starving: 饥饿的
				WaiterShift：等待的数量

		sema: 信号量、信号锁 => 最大并发数量 => 初始为0
				uint8 => semaRoot
										=> 对应的锁对象 Mutex
                    => 平衡树 => 正在等待运行得 Goroutine
										=> 等待的数量

	sema: 数量（并发）=>
				1. 当一个 Goroutine 获取到锁， sema - 1
				2. 当一个 Goroutine 释放锁了，sema + 1
				3. 当 sema 为 0, 携程系统就是休眠的状态
				（Goroutine 都会进入到等待队列）
				4. 只要有一个 Goroutine 释放了锁，sema 就会有一个并发量
				那么，系统就会从等待的队列里唤醒一个 Goroutine


*/
