package main

import (
	"context"
	"go-basics/11_grpc-advanced/basic1/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net"
)

const TOKEN = "abcdefg1234567"

type User struct {
}

func (u *User) GetUserInfo(
	ctx context.Context,
	req *proto.UserInfoRequest,
) (
	*proto.UserInfoResponse,
	error,
) {

	md, _ := metadata.FromIncomingContext(ctx)

	clientToken := md.Get("Token")[0]

	// if req.Token != TOKEN {
	if clientToken != TOKEN {
		return &proto.UserInfoResponse{
			Code: 1,
			Msg:  "Invalid Token",
		}, nil
	}

	return &proto.UserInfoResponse{
		Code: 0,
		Msg:  "ok",
		Data: &proto.UserInfo{
			Uid:      1,
			Username: "xulei",
			CourseList: []proto.Course{
				proto.Course_GO,
				proto.Course_JAVA,
				proto.Course_RUST,
			},
		},
	}, nil
}

func main() {
	gServer := grpc.NewServer()

	proto.RegisterUserServer(gServer, &User{})
	listener, _ := net.Listen("tcp", ":8080")

	_ = gServer.Serve(listener)
}
