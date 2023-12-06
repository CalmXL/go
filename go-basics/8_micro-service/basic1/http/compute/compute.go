package main

import (
	"encoding/json"
	"net/http"
)

type RequestBody struct {
	Type string `json:"type"`
	A    int    `json:"a"`
	B    int    `json:"b"`
}

type ResponseBody struct {
	A   int `json:"a"`
	B   int `json:"b"`
	Res int `json:"res"`
}

func main() {
	http.HandleFunc("/compute", func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			rw.WriteHeader(405) // Method Not allowed
			return
		}

		/**
		  r.Body => 解码 => requestBody
		*/
		requestBody := &RequestBody{}
		_ = json.NewDecoder(r.Body).Decode(requestBody)

		res := 0

		switch requestBody.Type {
		case "PLUS":
			res = plus(requestBody.A, requestBody.B)
		case "MINUS":
			res = minus(requestBody.A, requestBody.B)
		}

		rw.Header().Set("Content-Type", "application/json")

		jsonData, _ := json.Marshal(map[string]any{
			"msg":  "ok",
			"code": 0,
			"data": ResponseBody{
				A:   requestBody.A,
				B:   requestBody.B,
				Res: res,
			},
		})

		_, _ = rw.Write(jsonData)
	})

	_ = http.ListenAndServe(":8080", nil)
}

func plus(a int, b int) int {
	return a + b
}

func minus(a int, b int) int {
	return a - b
}
