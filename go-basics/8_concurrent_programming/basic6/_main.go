package main

import (
	"fmt"
	"time"
)

/**
  什么是上下文？
		上下文是程序运行得上下文，也就是在程序运行得时候，所有子孙程序都能通过访问到一个对象，
这个对象里有着所有子孙程序都需要的或者部分需要的数据或者功能，那么这个对象就是这个程序的上下文。

		Go 的 Context 本质上与我们讲的上下文概念是相同的
		但是 Go 的 Context 内部提供的是传递键值对
		超时取消功能
*/

type PrivateInfo struct {
	name string
	age  uint
}

type Context struct {
	info PrivateInfo
}

func (ctx *Context) getName() string {
	return ctx.info.name
}

func (ctx *Context) getAge() uint {
	return ctx.info.age
}

func thGetStuInfo(ctx Context) {
	fmt.Println(ctx.getName())
	fmt.Println(ctx.getAge())
	// fmt.Printf("距离25岁之前还有 %d 岁", 25-ctx.getAge())
	go test(ctx)
}

func stuGetThInfo(ctx Context) {
	fmt.Println(ctx.getName())
	fmt.Println(ctx.getAge())
}

func test(ctx Context) {
	// fmt.Println(25 - ctx.getAge())
	fmt.Printf("距离25岁之前还有 %d 岁", 25-ctx.getAge())
}

func main() {
	
	stuCtx := Context{
		info: PrivateInfo{
			name: "xulei",
			age:  16,
		},
	}
	
	thCtx := Context{
		info: PrivateInfo{
			name: "xuhong",
			age:  18,
		},
	}
	
	go thGetStuInfo(stuCtx)
	go stuGetThInfo(thCtx)
	// go test(stuCtx)
	
	time.Sleep(time.Second * 3)
}
