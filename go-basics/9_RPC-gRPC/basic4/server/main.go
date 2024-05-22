package main

import (
	"context"
	proto "go-basics/9_RPC-gRPC/basic4/proxy"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
}

func (s *Server) Calculate(
	ctx context.Context,
	req *proto.CalcRequest,
) (*proto.CalcResponse, error) {
	var result int32 = 0

	switch req.Type {
	case 0:
		result = req.Num1 + req.Num2
	case 1:
		result = req.Num1 - req.Num2
	case 2:
		result = req.Num1 * req.Num2
	case 3:
		result = req.Num1 / req.Num2
	}

	return &proto.CalcResponse{
		Num1:   req.Num1,
		Num2:   req.Num2,
		Type:   req.Type,
		Result: result,
	}, nil
}

func main() {
	gServer := grpc.NewServer()

	// 注册服务
	proto.RegisterCalculatorServer(gServer, &Server{})

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err.Error())
	}

	err = gServer.Serve(listener)
	if err != nil {
		panic(err.Error())

	}
}
