package main

import (
	"context"
	"go-basics/10_protocol-buffer/basic6/data"
	"go-basics/10_protocol-buffer/basic6/proto"
	"google.golang.org/grpc"
	"net"
)

type Student struct{}

func (s *Student) GetStudent(
	ctx context.Context,
	req *proto.StudentRequest,
) (
	*proto.StudentResponse,
	error,
) {
	var res *proto.StudentResponse

	for i := 0; i < len(data.Students); i++ {
		if data.Students[i].Sid == req.Sid {
			res = data.Students[i]
		}
	}

	return res, nil
}

func main() {
	gServer := grpc.NewServer()

	proto.RegisterStudentServer(gServer, &Student{})

	listener, _ := net.Listen("tcp", ":8080")

	_ = gServer.Serve(listener)
}
