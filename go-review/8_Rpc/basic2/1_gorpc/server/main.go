package main

import (
	"fmt"
	"net"
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

type ComputeService struct{}

func (cs *ComputeService) Add(req RequestBody, resp *ResponseBody) error {
	A := req.A
	B := req.B
	
	*resp = ResponseBody{
		Code: 0,
		Msg:  "ok",
		Data: ResponseData{
			A:      A,
			B:      B,
			Type:   "+",
			Result: A + B,
		},
	}
	
	return nil
}

// func (cs *ComputeService) Minus(req RequestBody, resp *ResponseBody) error {
// 	A := req.A
// 	B := req.B
//
// 	*resp = ResponseBody{
// 		Code: 0,
// 		Msg:  "ok",
// 		Data: ResponseData{
// 			A:      A,
// 			B:      B,
// 			Type:   "-",
// 			Result: A - B,
// 		},
// 	}
//
// 	return nil
// }

func (c *ComputeService) Minus(req RequestBody, resp *ResponseBody) error {
	
	A := req.A
	B := req.B
	
	fmt.Println(A, B)
	
	*resp = ResponseBody{
		Code: 0,
		Msg:  "ok",
		Data: ResponseData{
			A:      A,
			B:      B,
			Type:   "-",
			Result: A - B,
		},
	}
	
	return nil
}

func main() {
	// 1. 注册 rpc 服务
	_ = rpc.RegisterName("ComputeService", &ComputeService{})
	// 2. 创建网络监听器
	listener, _ := net.Listen("tcp", ":3000")
	
	for {
		// 3. 接收客户端来的连接
		conn, _ := listener.Accept()
		
		// 4. 给客户端提供 rpc 服务
		go rpc.ServeConn(conn)
	}
	
}
