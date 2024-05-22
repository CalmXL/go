package main

import (
	"go-basics/8_micro-service/basic2/6_new-gorpc/proxy/server"
	"go-basics/8_micro-service/basic2/6_new-gorpc/proxy/service"
	"net"
	"net/rpc"
)

/**
class Compute {
	plus() {}
}
*/

func main() {

	listener, _ := net.Listen("tcp", ":8080")
	// _ = rpc.RegisterName("ComputeService", &service.ComputeService{})

	_ = server.RegisterComputeService(&service.ComputeService{})

	for {
		conn, _ := listener.Accept()
		go rpc.ServeConn(conn)
	}
}
