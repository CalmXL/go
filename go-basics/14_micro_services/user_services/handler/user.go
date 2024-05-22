package handler

import (
	"14_micro_services/user_services/config"
	"14_micro_services/user_services/model"
	"14_micro_services/user_services/proto"
	"14_micro_services/user_services/util"
	"context"
	"fmt"
	"regexp"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
	proto.UnimplementedUserServer
}

func UserModelToResponse(user model.User) *proto.UserInfo {
	return &proto.UserInfo{
		Id:           user.ID,
		MobileNumber: user.MobileNumber,
		NickName:     user.NickName,
		Password:     user.Password,
		Gender:       user.Gender,
		Role:         user.Role,
	}
}

func (u *User) GetUserList(c context.Context, r *proto.PageInfo) (*proto.UserList, error) {
	var userList []model.User
	result := u.DB.Find(&userList)

	if result.Error != nil {
		return nil, result.Error
	}

	resp := &proto.UserList{}
	resp.UserCount = int32(result.RowsAffected)

	u.DB.Scopes(util.Paginate(r.Page, r.PageSize)).Find(&userList)

	for _, user := range userList {
		userInfo := UserModelToResponse(user)
		resp.Users = append(resp.Users, userInfo)
	}

	return resp, nil
}
func (u *User) GetUser(c context.Context, r *proto.UserInfo) (*proto.UserInfo, error) {

	/*
		id
		mobileNumber
	*/

	var user *model.User
	var err error

	if r.Id > 0 {
		user, err = getUserByWhatever(u.DB, r, "ID")
	}

	if r.MobileNumber != "" {
		user, err = getUserByWhatever(u.DB, r, "MobileNumber")
	}

	if err != nil {
		return nil, err
	}

	return UserModelToResponse(*user), nil
}
func (u *User) CreateUser(c context.Context, r *proto.UserInfo) (*proto.UserInfo, error) {

	/*
		MobileNumber unique
		password
		Nickname
	*/

	var user model.User

	result := u.DB.Where(&model.User{
		MobileNumber: r.MobileNumber,
	}).First(&user)

	// if result.Error != nil {
	// 	return nil, status.Errorf(codes.Internal, "Query  Failed")
	// }

	if result.RowsAffected != 0 {
		return nil, status.Errorf(codes.AlreadyExists, "User exists")
	}

	user.MobileNumber = r.MobileNumber
	user.NickName = r.NickName
	user.Password = util.GeneratePassword(r.Password)

	result = u.DB.Create(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}

	resp := UserModelToResponse(user)

	return resp, nil
}
func (u *User) UpdateUser(c context.Context, r *proto.UserInfo) (*proto.UserInfo, error) {

	user, err := getUserByWhatever(u.DB, r, "ID")
	if err != nil {
		return nil, err
	}

	user.NickName = r.NickName
	user.Gender = r.Gender
	user.MobileNumber = r.MobileNumber

	result := u.DB.Save(&user)

	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}

	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	return UserModelToResponse(*user), nil
}
func (u *User) VerifyPassword(c context.Context, r *proto.PasswordVerify) (*proto.PasswordVeryifyPass, error) {

	/*
		id => database => encodedPassword
		rawPassword
	*/

	var user model.User

	result := u.DB.First(&user, r.Id)

	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}

	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	encodedPassword := user.Password
	isPass := util.VerifyPassword(r.RawPassword, encodedPassword)

	return &proto.PasswordVeryifyPass{
		IsPass: isPass,
	}, nil
}

func (u *User) UpdateMobileNumber(c context.Context, r *proto.UserInfo) (*proto.UserInfo, error) {
	user, err := getUserByWhatever(u.DB, r, "ID")

	if err != nil {
		return nil, err
	}

	matched, err := regexp.MatchString(config.REGPhoneNumber, r.MobileNumber)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if !matched {
		return nil, status.Errorf(codes.InvalidArgument, "MobileNumber invalid")
	}

	user.MobileNumber = r.MobileNumber
	result := u.DB.Save(&user)

	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}

	return UserModelToResponse(*user), nil
}
func (u *User) UpdatePassword(c context.Context, r *proto.UserInfo) (*proto.UserInfo, error) {
	user, err := getUserByWhatever(u.DB, r, "ID")

	if err != nil {
		return nil, err
	}

	user.Password = util.GeneratePassword(r.Password)
	result := u.DB.Save(&user)

	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}

	return UserModelToResponse(*user), nil

}

func getUserByWhatever(db *gorm.DB, r *proto.UserInfo, filed string) (*model.User, error) {
	var user model.User

	fmt.Println(filed)

	switch filed {
	case "ID":
		user = model.User{
			ID: r.Id,
		}
	case "MobileNumber":
		user = model.User{
			MobileNumber: r.MobileNumber,
		}
	default:
		user = model.User{
			ID: r.Id,
		}
	}

	result := db.Where(&user).First(&user)

	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "Query  Failed")
	}

	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	return &user, nil
}
