package main

import (
	"14_micro_services/user_services/config"
	"14_micro_services/user_services/proto"
	"14_micro_services/user_services/util"
	"context"
	"fmt"
	"log"
)

func main() {
	db, errConn := util.DBConnect(config.DBName)
	if errConn != nil {
		log.Fatalf("Failed to connect database: %v", errConn)
		return
	}

	conn, err := util.GrpcDial(*config.IP, *config.Port)
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	client := proto.NewUserClient(conn)

	// 获取分页列表
	// resp, _ := client.GetUserList(context.Background(), &proto.PageInfo{
	// 	Page:     2,
	// 	PageSize: 3,
	// })

	// fmt.Println(resp)

	// for _, user := range resp.Users {
	// 	fmt.Println(user)
	// }

	// 查询单个用户
	resp, _ := client.GetUser(context.Background(), &proto.UserInfo{
		Id: 1,
		// MobileNumber: "13800000011",
	})

	fmt.Println(resp)

	// 创建单个用户
	// resp, _ := client.CreateUser(context.Background(), &proto.UserInfo{
	// 	NickName:     "xiaohong",
	// 	MobileNumber: "15663399210",
	// 	Password:     "xL1210...",
	// })
	//
	// fmt.Println(resp)

	// 更新用户
	// resp, _ := client.UpdateUser(context.Background(), &proto.UserInfo{
	// 	Id:           30,
	// 	NickName:     "xiaoxu",
	// 	MobileNumber: "1000000",
	// })
	//
	// fmt.Println(resp)

	// 验证密码
	// resp, err := client.VerifyPassword(context.Background(), &proto.PasswordVerify{
	// 	Id:          29,
	// 	RawPassword: "a87654321",
	// })
	//
	// fmt.Println(resp.IsPass)

	// 更新 mobilePhone
	// resp, _ := client.UpdateMobileNumber(context.Background(), &proto.UserInfo{
	// 	Id:           1,
	// 	MobileNumber: "15663399210",
	// })
	//
	// fmt.Println(resp)

	// 更新 password
	// resp, _ := client.UpdatePassword(context.Background(), &proto.UserInfo{
	// 	Id:       29,
	// 	Password: "a87654321",
	// })
	//
	// fmt.Println(resp)

	// 创建表
	// if errMigrate := db.AutoMigrate(
	// 	&model.User{},
	// ); errMigrate != nil {
	// 	log.Fatalf("Failed to migrate database: %v", errMigrate)
	// 	return
	// }

	// 添加随机人员
	// generateUsers(db)
	defer util.DBClose(db)
}
