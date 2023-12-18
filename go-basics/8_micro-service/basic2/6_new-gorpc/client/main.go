package main

import (
	"fmt"
	"go-basics/8_micro-service/basic2/6_new-gorpc/proxy/client"
	"go-basics/8_micro-service/basic2/6_new-gorpc/proxy/service"
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

	args := service.RequestBody{
		A: 300,
		B: 200,
	}

	reply := &service.ResponseBody{}
	// dial, _ := rpc.Dial("tcp", ":8080")

	computeService := client.NewComputeServiceClient("tcp", ":8080")
	err := computeService.Minus(args, reply)
	if err != nil {
		return
	}
	// _ = dial.Call("ComputeService.Minus", args, reply)
	// // compute.Minus(args, reply)
	fmt.Println(reply)

}
