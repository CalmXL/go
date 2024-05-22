package main

import (
	"fmt"
	"net/rpc"
)

type RequestBody struct {
	A int `json:"a"`
	B int `json:"b"`
}

type ResponseData struct {
	A      int    `json:"a"`
	B      int    `json:"b"`
	Type   string `json:"type"`
	Result int    `json:"result"`
}

type ResponseBody struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data ResponseData `json:"data"`
}

func main() {

	args := RequestBody{
		A: 300,
		B: 200,
	}

	reply := &ResponseBody{}

	/**
	  1.使用 rpc 服务拨号

			参数：
				1. 网络协议
				2. 拨号地址
	*/
	dial, _ := rpc.Dial("tcp", ":8080")
	/**
	  2. 利用拨号实例进行远程方法调用
			参数
				1. 什么服务中的什么方法（必须使用字符串）
				2. req RequestBody 请求参数
				3. resp *ResponseBody 结果响应指针
	*/
	_ = dial.Call("ComputeService.Minus", args, reply)

	fmt.Println(reply)

	/**
	  tcp => 传递的是二进制的字节流
		go => 一定是把数据转换成二进制字节流，用 TCP 进行传输

		数据 => 二进制 => 服务端
		服务端 => 二进制 => 可读取数据 => 程序运行 => 打包数据 => 二进制 => 客户端
		客户端 => 解包数据 => 可读取的数据

	  打包: Go 可读取数据 转换为二进制字节流的过程 => 序列化  编码的过程
		解包: 二进制字节流转换为 Go 可读数据的过程 => 反序列化  解码的过程

		编解码的过程是需要约定的 => 序列化协议

		常见的协议:
		JSON
		XML
		gob  它是 Go 语言中达成的序列化协议
	       性能好
	*/
}
