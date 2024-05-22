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

type Compute struct{}

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
	registerServices("Compute", &Compute{}, "tcp", "5173")
}

func registerServices(name string, service any, portcol string, port string) {
	_ = rpc.RegisterName(name, service)
	listeren, _ := net.Listen(portcol, ":"+port)
	
	for {
		connect, _ := listeren.Accept()
		
		go rpc.ServeConn(connect)
	}
}

func plus(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}
