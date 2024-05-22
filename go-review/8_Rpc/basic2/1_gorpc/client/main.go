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
		A: 100,
		B: 200,
	}
	
	reply := &ResponseBody{}
	
	// 1. 使用 rpc 服务拨号
	dial, _ := rpc.Dial("tcp", ":3000")
	// 2. 利用拨号实例
	_ = dial.Call("ComputeService.Minus", args, reply)
	
	fmt.Println(reply)
}
