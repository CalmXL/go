package client

import "net/rpc"

/**
  dial, _ := rpc.Dial("tcp", ":8080")

	_ = dial.Call("ComputeService.Minus", args, reply)
	// compute.Minus(args, reply)
	fmt.Println(reply)
*/

type ComputeStub struct {
	*rpc.Client
}

func (c *ComputeStub) Add(req, res) error {
	err := c.Call(service.ComputeName+".Add".req, res)
	if err != nil {
		return err
	}
}
