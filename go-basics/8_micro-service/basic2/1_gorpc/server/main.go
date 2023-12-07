package main

import (
	"fmt"
	"net"
	"net/rpc"
)

/**
class Compute {
	plus() {}
}
*/

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

// rpc服务

type ComputeService struct {
}

func (c *ComputeService) Add(req RequestBody, resp *ResponseBody) error {

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

	// 只要在这里程序发生了错误，可以直接返回错误
	// 如果没有错误必须返回 nil
	return nil
}

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

	// 只要在这里程序发生了错误，可以直接返回错误
	// 如果没有错误必须返回 nil
	return nil
}

func (c *ComputeService) Multiply(req RequestBody, resp *ResponseBody) error {

	A := req.A
	B := req.B

	*resp = ResponseBody{
		Code: 0,
		Msg:  "ok",
		Data: ResponseData{
			A:      A,
			B:      B,
			Type:   "*",
			Result: A * B,
		},
	}

	// 只要在这里程序发生了错误，可以直接返回错误
	// 如果没有错误必须返回 nil
	return nil
}

func (c *ComputeService) Div(req RequestBody, resp *ResponseBody) error {

	A := req.A
	B := req.B

	*resp = ResponseBody{
		Code: 0,
		Msg:  "ok",
		Data: ResponseData{
			A:      A,
			B:      B,
			Type:   "/",
			Result: A / B,
		},
	}

	// 只要在这里程序发生了错误，可以直接返回错误
	// 如果没有错误必须返回 nil
	return nil
}

func main() {

	// 1. 注册 rpc 服务
	// ComputeService => Rpc 需要这个名字的服务注册到 rpc 中
	/**
	  1. 注册 rpc 服务名称
		2. Rpc 服务结构体引用注册
	*/

	_ = rpc.RegisterName("ComputeService", &ComputeService{})

	// 2. 创建网络监听器
	/**
	  1. 网络协议
	  2. 服务地址（主要端口号）
	*/
	listener, _ := net.Listen("tcp", ":8080")

	for {
		/**
		  3. 接收客户端来的连接
				 就是为了获取连接实例
		     有了连接实例 RPC才能对该连接提供服务
		*/
		conn, _ := listener.Accept()

		/**
		  4.给客户端连接提供 RPC 服务
		*/
		go rpc.ServeConn(conn)
	}
}
