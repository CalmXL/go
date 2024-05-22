package main

import (
	"net"
	"net/rpc"
)

type RequestData struct {
	Type string `json:"type"`
	A    int    `json:"a"`
	B    int    `json:"b"`
}

type ResponseData struct {
	A   int `json:"a"`
	B   int `json:"b"`
	Res int `json:"res"`
}

type ResponseBody struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data ResponseData `json:"data"`
}

type Compute struct {
}

// resp 地址 => resp 的指针，往对应的空间放东西
// rpc 服务的方法 必须返回 error

func (cm *Compute) Do(req RequestData, resp *ResponseBody) error {
	A := req.A
	B := req.B

	*resp = ResponseBody{
		Code: 0,
		Msg:  "ok",
		Data: ResponseData{
			A:   A,
			B:   B,
			Res: 0,
		},
	}

	switch req.Type {
	case "PLUS":
		resp.Data.Res = plus(A, B)
	case "MINUS":
		resp.Data.Res = minus(A, B)
	}

	return nil
}

func main() {
	registerServices("Compute", &Compute{}, "tcp", "8080")
}

func registerServices(name string, service any, protocol, port string) {
	_ = rpc.RegisterName(name, service)
	listener, _ := net.Listen(protocol, ":"+port)

	for {
		connect, _ := listener.Accept()
		// 创建连接
		go rpc.ServeConn(connect)
	}
}

func plus(a int, b int) int {
	return a + b
}

func minus(a int, b int) int {
	return a - b
}
