package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go-basics/10_protocol-buffer/basic5/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, _ := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	client := proto.NewCourseClient(conn)

	computerCoureses, _ := client.GetCourseList(context.Background(), &proto.CourseRequest{
		Type: "COMPUTER",
	})

	languageCourses, _ := client.GetCourseList(context.Background(), &proto.CourseRequest{
		Type: "LANGUAGE",
	})

	jsonComputerData, _ := json.Marshal(computerCoureses)
	jsonLanguageData, _ := json.Marshal(languageCourses)

	fmt.Println(string(jsonComputerData))
	fmt.Println(string(jsonLanguageData))
}
