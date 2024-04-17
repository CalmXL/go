package main

import "fmt"

type User interface {
	login()
	register()
	error(string)
}

type Admin interface {
	User
	adminLogin()
}

type SystemAdmin struct {
	username   string
	password   string
	passcode   string
	randomCode string
}

func (s SystemAdmin) login() {
	usernameLen := len(s.username)
	passwordLen := len(s.password)
	
	if usernameLen < 6 || usernameLen > 20 {
		s.error("用户名长度不正确")
		return
	}
	
	if passwordLen < 8 || passwordLen > 20 {
		s.error("密码长度不正确")
		return
	}
}

func (s SystemAdmin) register() {
	// TODO implement me
	panic("implement me")
}

func (s SystemAdmin) error(s2 string) {
	fmt.Println(s2)
}

func (s SystemAdmin) adminLogin() {
	// TODO implement me
	panic("implement me")
}

func main() {
	
	var admin Admin = SystemAdmin{
		username:   "xulei111",
		password:   "123456dasdasda",
		passcode:   "123456",
		randomCode: "893812",
	}
	
	admin.login()
}
