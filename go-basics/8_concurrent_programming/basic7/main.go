package main

import (
	"context"
	"fmt"
)

func main() {
	// 1. 创建空的上下文
	// 创建一个需要派生子上下文的根
	ctx1 := context.Background()
	// 创建一个不明确的上下文
	// ctx2 := context.TODO()

	fmt.Println(ctx1)

	/**
	    type Context interface {
			1. Deadline 返回 context 被取消的时间， 截止日期
			Deadline() (deadline time.Time, ok bool)
			2. 返回一个信号 Channel, 提示 context 被取消
				Goroutine 就可以在 Done() 分支接收到信号时，中断运行
			Done() <-chan struct{}
			3. 返回 context 被取消的原因 (你做了什么操作导致的被取消)
				Done() 分支 接收到信号
				a. 如果 context 被取消，会返回 canceled 错误
				b. context 超时，会返回 DeadlineExceeded 错误
			Err() error
			4. 可以从 context 中获取 key 对应的值
	            基本上不会通过 value 传递参数之类的数据
				token, id
			Value(key any) any
		}

		func Background() Context {
			return backgroundCtx{}
		}
	*/

	/**
		  cancelCtx := context.withCancel(ctx1)
			case <- Done()

	        valueCtx := context.withValue(ctx1)
			valueCtx.Value()

			timeoutCtx := context.withTimeout(cancelCtx, 2 * time.Second)

			deadlineCtx := context.withDeadline(ctx1, time.Now().Add(2 * time.Second))
	*/
}
