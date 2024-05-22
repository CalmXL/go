package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
)

type PrivateInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	privateInfo := PrivateInfo{
		Name: "xulei",
		Age:  27,
	}

	infoBuf := encode(privateInfo)
	data := decode(infoBuf)

	fmt.Println(data)
	// fmt.Println(io.ReadAll(infoBuf))
}

func encode(data PrivateInfo) *bytes.Buffer {
	// Buffer 缓冲区 => 传输数据的时候需要一个临时存放数据的缓冲区
	var buf bytes.Buffer

	// 将字节流缓冲区包装成一个编码器
	encoder := gob.NewEncoder(&buf)
	// 对数据进行编码，并对编码的结果装入 Buffer 中
	_ = encoder.Encode(data)

	return &buf
}

func decode(data any) *PrivateInfo {
	// 声明一个解码后的数据容器
	var result PrivateInfo

	// 将传入的数据转换成可读的类型
	// 要对数据进行包装，数据必须是可读类型数据 io.Reader
	dt := data.(io.Reader)

	// 将 io.Reader 类型的数据包装成一个解码器
	decoder := gob.NewDecoder(dt)
	_ = decoder.Decode(&result)

	return &result
}
