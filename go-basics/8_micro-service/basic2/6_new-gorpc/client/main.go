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
	dial, _ := rpc.Dial("tcp", ":8080")

	_ = dial.Call("ComputeService.Minus", args, reply)
	// compute.Minus(args, reply)
	fmt.Println(reply)

}
