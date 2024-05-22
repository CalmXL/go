package main

import (
	"fmt"
	"time"
)

var timeout1 = make(chan bool)
var timeout2 = make(chan bool)
var timeout3 = make(chan bool)

func g1() {
	defer close(timeout1)
	go g2()
	
	for {
		select {
		case <-timeout1:
			fmt.Println("请求1 超时！")
			fmt.Println("请求1 已中断")
			return
		default:
			fmt.Println("请求中...1")
			time.Sleep(time.Second)
		}
	}
}

func g2() {
	defer close(timeout2)
	go g3()
	for {
		select {
		case <-timeout2:
			fmt.Println("请求2 超时！")
			fmt.Println("请求2 已中断")
			return
		default:
			fmt.Println("请求中...2")
			time.Sleep(time.Second)
		}
	}
}

func g3() {
	defer close(timeout3)
	for {
		select {
		case <-timeout3:
			fmt.Println("请求3 超时！")
			fmt.Println("请求3 已中断")
			return
		default:
			fmt.Println("请求中...3")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	go g1()
	
	time.Sleep(time.Second * 3)
	timeout1 <- true
	timeout2 <- true
	timeout3 <- true
	time.Sleep(time.Second)
}
