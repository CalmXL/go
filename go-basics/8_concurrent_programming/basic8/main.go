package main

import (
	"context"
	"fmt"
	"time"
)

func req(ctx1 context.Context, ctx2 context.Context) {
	// tick := time.Tick(1 * time.Second)
	//
	// for range tick {
	// 	select {
	// 	case <-ctx1.Done():
	// 		fmt.Println("请求已中断")
	// 		return
	// 	default:
	// 		fmt.Println("请求中...")
	// 	}
	// }

	select {
	case <-ctx1.Done():
		fmt.Println("请求已中断")
		return
	default:
		// fmt.Println("请求中...")
		go res(ctx2)
	}

}

func res(ctx context.Context) {
	tick := time.Tick(1 * time.Second)

	for range tick {
		select {
		case <-ctx.Done():
			fmt.Println("响应已中断")
			fmt.Println(ctx.Err())
			return
		default:
			val := ctx.Value("value")
			fmt.Println(val)
		}
	}
}

func main() {

	emptyCtx := context.Background()

	/**
	  参数1: 父上下文
	  返回值:
	        1. 返回 withCancel 派生的上下文
			2. cancel 方法，调用会取消上下文
				通过 ctx.Done() 接收上下文取消信号
				从而在 sleect 分支中进行 Goroutine 的运行终止
	*/
	// cancelCtx, cancel := context.WithCancel(emptyCtx)
	// fmt.Println(cancelCtx, cancel)

	// timeoutCtx, cancel := context.WithTimeout(emptyCtx, 5*time.Second)

	deadlineCtx, _ := context.WithDeadline(emptyCtx, time.Now().Add(5*time.Second))

	/**
	  参数1： 父上下文
		参数2：key
		参数3: value
		返回值: 通过 WithValue 返回 withCancel 派生的上下文
	*/
	valueCtx := context.WithValue(deadlineCtx, "value", 1)
	// fmt.Println(valueCtx)

	go req(deadlineCtx, valueCtx)

	// time.Sleep(3 * time.Second)

	// cancel()

	time.Sleep(10 * time.Second)

	/**
	1. go req
	2. 五秒 => req default
	3. cancel => req (x 监听不到) => res => ctx.Done()
	4.
	*/
}
