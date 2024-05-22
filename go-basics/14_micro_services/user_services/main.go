package main

import (
	"14_micro_services/user_services/config"
	"14_micro_services/user_services/handler"
	"14_micro_services/user_services/proto"
	"14_micro_services/user_services/util"
	"log"
)

func main() {
	db, err := util.DBConnect(config.DBName)

	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
		return
	}

	gServer, listener, err := util.GrpcServer(*config.IP, *config.Port, nil)

	if err != nil {
		log.Fatalf("Failed to create gRpc server: %v", err)
		return
	}

	// 注册某一个 gRpc 的 server
	proto.RegisterUserServer(gServer, &handler.User{
		DB: db,
	})

	if err := gServer.Serve(listener); err != nil {
		log.Fatalf("Failed to start gRpc server: %v", err)
		return
	}
}
