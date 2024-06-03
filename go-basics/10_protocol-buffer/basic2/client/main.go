package main

import (
	"context"
	"fmt"
	"go-basics/10_protocol-buffer/basic2/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	/**
	  1. 向服务端拨号
		2. 实例化一个 gRpc 客户端 => GetData
		3. 调用客户端方法
		4. 取响应信息
	*/

	conn, _ := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	client := proto.NewMyInfoClient(conn)

	data, err := client.GetData(context.Background(), &proto.MyInfoRequest{
		Name:       "xulei",
		Age:        18,
		IsMarriage: false,
	})
	if err != nil {
		return
	}

	fmt.Println(data)
}
