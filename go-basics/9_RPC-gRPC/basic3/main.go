package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	__ "go-basics/9_RPC-gRPC/basic3/proto"
)

type Person struct {
	Id         int32  `json:"id"`
	Name       string `json:"name"`
	IsMarriage bool   `json:"is_marriage"`
}

func main() {
	phInfo := __.Person{
		Id:         1,
		Name:       "xiaoyesensen",
		IsMarriage: true,
	}

	jsonInfo := Person{
		Id:         1,
		Name:       "xiaoyesensen",
		IsMarriage: true,
	}

	phResp, _ := proto.Marshal(&phInfo)
	jsonResp, _ := json.Marshal(jsonInfo)
	//
	// fmt.Println(phResp)
	// fmt.Println(jsonResp)
	// [8 1 18 12 120 105 97 111 121 101 115 101 110 115 101 110 24 1]
	// [123 34 105 100 34 58 49 44 34 110 97 109 101 34 58 34 120 105 97 111 121 101 115 101 110 115 101 110 34 44 34 105 115 95 109 97 114 114 105 97 103 101 34 58 116 114 117 101 125]

	fmt.Println(string(phResp), string(jsonResp))
	//                                            x
	// iaoyesensen {"id":1,"name":"xiaoyesensen","is_marriage":true}

	fmt.Println(len(phResp), len(jsonResp)) // 18 49
}
