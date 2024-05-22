package util

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GrpcServer(ip, port string, options []grpc.ServerOption) (*grpc.Server, net.Listener, error) {

	// gServer, listener, error
	gServer := grpc.NewServer(options...)
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", ip, port))

	return gServer, listener, err
}

func GrpcDial(ip, port string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", ip, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	return conn, err
}
