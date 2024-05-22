package main

import (
	"fmt"
	"io"
	"net/http"
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

type Connect struct {
	io.Writer
	io.ReadCloser
}

func main() {
	_ = rpc.RegisterName("ComputeService", &ComputeService{})

	http.HandleFunc("/compute", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = Connect{
			ReadCloser: r.Body,
			Writer:     w,
		}

		_ = rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	_ = http.ListenAndServe(":8080", nil)
}
