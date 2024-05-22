package client

import (
	"fmt"
	"go-basics/9_RPC-gRPC/basic2/6_new-gorpc/proxy/service"
	"net/rpc"
)

/**
  dial, _ := rpc.Dial("tcp", ":8080")

	_ = dial.Call("ComputeService.Minus", args, reply)
	// compute.Minus(args, reply)
	fmt.Println(reply)
*/

// class ComputeStub {
// 	NewComputeServiceClient() {
// 		conn, err := rpc.Dial(protocol, address)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		return new ComputeStub(conn)
// 	}
//
// 	Add (req, res) {
// err := c.Call(service.ComputeName+".Add", req, res)
// 	}
// }

type ComputeStub struct {
	*rpc.Client
}

func NewComputeServiceClient(protocol, address string) ComputeStub {
	conn, err := rpc.Dial(protocol, address)
	if err != nil {
		panic(err.Error())
	}
	return ComputeStub{conn}
}

func (c *ComputeStub) Add(req service.RequestBody, res *service.ResponseBody) error {
	err := c.Call(service.ComputeName+".Add", req, res)
	if err != nil {
		return err
	}

	return nil
}

func (c *ComputeStub) Minus(req service.RequestBody, res *service.ResponseBody) error {
	fmt.Println("Minus")
	fmt.Println(service.ComputeName + ".Minus")
	err := c.Call(service.ComputeName+".Minus", req, res)
	fmt.Println(err)
	if err != nil {
		return err
	}

	return nil
}

func (c *ComputeStub) Multiply(req service.RequestBody, res *service.ResponseBody) error {
	err := c.Call(service.ComputeName+".Multiply", req, res)
	if err != nil {
		return err
	}

	return nil
}

func (c *ComputeStub) Division(req service.RequestBody, res *service.ResponseBody) error {
	err := c.Call(service.ComputeName+".Division", req, res)
	if err != nil {
		return err
	}

	return nil
}
