package main

import (
	"fmt"
	"go-basics/9_RPC-gRPC/basic2/6_new-gorpc/proxy/client"
	"go-basics/9_RPC-gRPC/basic2/6_new-gorpc/proxy/service"
)

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
