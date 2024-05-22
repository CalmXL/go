package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
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
	// 不需要用 rpc 进行拨号。因为 rpc 拨号返回的是拨号功能实例
	// 我们不需要使用拨号功能实例进行远程调用
	// 我们需要使用 client 实例进行远程调用
	// 而 net 拨号会直接返回一个客户端连接
	conn, _ := net.Dial("tcp", ":8080")
	// 指定客户端数据编码器
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	_ = client.Call("ComputeService.Minus", args, reply)

	fmt.Println(reply)
}
