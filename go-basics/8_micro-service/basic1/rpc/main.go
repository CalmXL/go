package main

import (
	"encoding/json"
	"net/http"
	"net/rpc"
)

type RequestBody struct {
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

func main() {
	http.HandleFunc("/plus", func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			rw.WriteHeader(405)
			return
		}

		requestBody := RequestBody{
			Type: "PLUS",
		}

		_ = json.NewDecoder(r.Body).Decode(&requestBody)
		responseBody := &ResponseBody{}

		applyService(
			"tcp",
			"localhost:8080",
			"Compute.Do",
			[]any{
				requestBody,
				responseBody,
			},
		)

		jsonStr, _ := json.Marshal(responseBody)

		_, _ = rw.Write(jsonStr)

	})

	http.HandleFunc("/minus", func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			rw.WriteHeader(405)
			return
		}

		requestBody := RequestBody{
			Type: "MINUS",
		}

		_ = json.NewDecoder(r.Body).Decode(&requestBody)
		responseBody := &ResponseBody{}

		applyService(
			"tcp",
			"localhost:8080",
			"Compute.Do",
			[]any{
				requestBody,
				responseBody,
			},
		)

		jsonStr, _ := json.Marshal(responseBody)

		_, _ = rw.Write(jsonStr)

	})

	_ = http.ListenAndServe(":3000", nil)
}

func applyService(protocol string, address string, serviceMethod string, args []any) {
	// 拨号
	dial, _ := rpc.Dial(protocol, address)
	err := dial.Call(serviceMethod, args[0], args[1])
	if err != nil {
		return
	}
}
