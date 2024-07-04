package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go-basics/10_protocol-buffer/basic6/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, _ := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	client := proto.NewStudentClient(conn)
	data, _ := client.GetStudent(context.Background(), &proto.StudentRequest{
		Sid: 2,
	})

	jsonData, _ := json.Marshal(data)

	fmt.Println(string(jsonData))
}
