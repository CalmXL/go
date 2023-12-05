package server

import (
	"net"
	"net/rpc"
)

func main() {
	// 1. 注册 RPC 服务

	/**
	  1. 注册 RPC 服务名称
		2. RPC 服务体结构引用注册
	*/
	// _ = rpc.RegisterName("rpc")

	// 2. 创建网络监听器
	/**
		1. 网络协议
	  2. 服务地址（主要是端口号）
	*/
	listener, _ := net.Listen("tcp", "8080")

	for {
		/**
		3. 接收客户端来的连接
				就是为了获取连接实例
				有了连接实例 RPC 才能对该连接提供服务
		*/
		conn, _ := listener.Accept()

		/**
		  给客户端连接提供 RPC 服务
		*/
		go rpc.ServeConn(conn)
	}
}
