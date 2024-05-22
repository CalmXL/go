package service

import "fmt"

const ComputeName = "ComputeService"

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
	fmt.Println("Minus")
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
