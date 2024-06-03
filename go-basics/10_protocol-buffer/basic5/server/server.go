package main

import (
	"context"
	"go-basics/10_protocol-buffer/basic5/data"
	"go-basics/10_protocol-buffer/basic5/proto"
	"google.golang.org/grpc"
	"net"
)

type Course struct {
}

func (c *Course) GetCourseList(ctx context.Context, req *proto.CourseRequest) (*proto.CourseResponse, error) {
	switch req.Type {
	case "LANGUAGE":
		return &proto.CourseResponse{
			Code: 0,
			Msg:  "ok",
			Data: data.LanguageCourse,
			Extra: map[string]string{
				"numberOfStudents": "300",
			},
		}, nil
	case "COMPUTER":
		return &proto.CourseResponse{
			Code: 0,
			Msg:  "ok",
			Data: data.ComputerCourse,
			Extra: map[string]string{
				"numberOfStudents": "300",
			},
		}, nil
	default:
		return &proto.CourseResponse{
			Code: 0,
			Msg:  "ok",
			Data: nil,
		}, nil
	}
}

func main() {
	// 1. 创建 grpc 服务
	gServer := grpc.NewServer()

	// 2. 注册
	proto.RegisterCourseServer(gServer, &Course{})
	listener, _ := net.Listen("tcp", ":8080")

	// 4. 启动
	_ = gServer.Serve(listener)
}
