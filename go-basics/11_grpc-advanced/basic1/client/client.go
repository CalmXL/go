package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go-basics/11_grpc-advanced/basic1/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

// var md = metadata.New(map[string]string{
// 	"token": "abcdefg1234567",
// })

var md = metadata.Pairs(
	"Token", "abcdefg1234567",
)

func main() {

	ctx := metadata.NewOutgoingContext(context.Background(), md)

	conn, _ := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	client := proto.NewUserClient(conn)

	// data, _ := client.GetUserInfo(context.Background(), &proto.UserInfoRequest{
	// 	Token: "xuleidashabi",
	// })

	data, _ := client.GetUserInfo(ctx, &proto.UserInfoRequest{})

	jsonData, _ := json.Marshal(data)

	fmt.Println(string(jsonData))
}
