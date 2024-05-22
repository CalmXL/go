package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RequestBody struct {
	Type string `json:"type"`
	A    int    `json:"a"`
	B    int    `json:"b"`
}

func main() {
	http.HandleFunc("/plus", func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			rw.WriteHeader(405)
			return
		}

		requestBody := &RequestBody{
			Type: "PLUS",
		}

		_ = json.NewDecoder(r.Body).Decode(requestBody)
		fmt.Println(requestBody)

		jsonStr, _ := json.Marshal(requestBody)

		resp, _ := http.Post(
			"http://localhost:5173/compute",
			"application/json",
			bytes.NewBuffer(jsonStr),
		)

		body, _ := io.ReadAll(resp.Body)

		_, _ = rw.Write(body)

	})

	http.HandleFunc("/minus", func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			rw.WriteHeader(405)
			return
		}

		requestBody := &RequestBody{
			Type: "MINUS",
		}

		_ = json.NewDecoder(r.Body).Decode(requestBody)

		jsonStr, _ := json.Marshal(requestBody)

		resp, _ := http.Post(
			"http://localhost:5173/compute",
			"application/json",
			bytes.NewBuffer(jsonStr),
		)

		body, _ := io.ReadAll(resp.Body)

		_, _ = rw.Write(body)

	})

	_ = http.ListenAndServe(":3000", nil)
}
