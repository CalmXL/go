package main

import (
	"context"
	"go-basics/10_protocol-buffer/basic2/proto"
	"google.golang.org/grpc"
	"net"
)

type MyInfo struct {
	proto.UnimplementedMyInfoServer
}

func (mi *MyInfo) GetData(ctx context.Context, req *proto.MyInfoRequest) (*proto.MyInfoResponse, error) {
	return &proto.MyInfoResponse{
		Name:       req.Name,
		Age:        req.Age,
		IsMarriage: req.IsMarriage,
	}, nil
}

func main() {
	/**
	  Server 流程
			1. 结构体
			2. 实现接口方法
			3. 创建 gRpc 服务
			4. 注册当前的服务
			5. 监听端口
			6. 启动 gRpc 服务
	*/

	gServer := grpc.NewServer()
	proto.RegisterMyInfoServer(gServer, &MyInfo{})
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		return
	}

	err = gServer.Serve(listener)
	if err != nil {
		return
	}
}
