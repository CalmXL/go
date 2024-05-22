package main

import (
	"context"
	"fmt"
	"time"
)

func g1(ctx context.Context) {
	go g2(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("请求1 已中断")
			return
		default:
			fmt.Println("请求中... 1")
			time.Sleep(time.Second)
		}
	}
}

func g2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("请求2 已中断")
			return
		default:
			fmt.Println("请求中... 2")
			time.Sleep(time.Second)
		}
	}
}

func g3(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("请求3 已中断")
			return
		default:
			fmt.Println("请求中... 3")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	/**
	  context.Background() => context 实例
		context.WithCancel => context 实例
	                         cancel方法 => 通过 Done 通道发送消息
	*/
	ctx, cancel := context.WithCancel(context.Background())

	go g1(ctx)

	time.Sleep(time.Second * 3)
	cancel()
	time.Sleep(time.Second)
}
