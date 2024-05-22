package main

import (
	"context"
	"fmt"
	proto "go-basics/9_RPC-gRPC/basic4/proxy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err.Error())
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			panic(err.Error())
		}
	}(conn)

	client := proto.NewCalculatorClient(conn)

	resp, err := client.Calculate(context.Background(), &proto.CalcRequest{
		Num1: 25,
		Num2: 4,
		Type: proto.CalcType_MULTIPLY,
	})

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(resp)
}
