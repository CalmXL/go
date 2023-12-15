package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
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
	_ = rpc.RegisterName("ComputeService", &ComputeService{})
	listener, _ := net.Listen("tcp", ":8080")

	for {
		conn, _ := listener.Accept()
		// 对编码协议进行指定 (方法指定)

		// 给客户端提供 rpc 服务
		// go rpc.ServeConn(conn) // 默认 gob 序列化协议

		// 可以指定编码器，且可以给客户端连接提供 rpc 服务
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
