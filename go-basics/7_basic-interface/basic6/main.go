package main

import "fmt"

/**
接口是拿来做什么的?
	接口是用来规范功能（方法）实现用的
	接口中保存着每个需要实现方法的定义规则

	type A interface {
		test1(string) string
		test2(int) int
	}

	接口主要可以用来创建鸭子类型（统一每个对象的类型，做到同一实现方法的调用）

  ts
	interface IObj {
		a: number,
		b: number,
	}
	const obj: IObj = {
		a: 1,
		b: 2
	}

	interface ITest {
		plus()
	}

	class Test implements ITest{
		a: number;
		b: number;

		constructor(a: number, b: number) {
			this.a = a;
			this.b = b;
		}

		plus() {
			return this.a + this.b
		}
	}

	const obj = new Test(1, 2)

	// --------------------
  Go

	type IObj interface {
		plus()
	}

	type Obj struct {
		a int
		b int
	}

	func (o *Obj) plus() {
		return o.a + o.b
	}

	obj := Obj{
		a: 1,
		b: 2,
	}

*/

/**
  interface 的嵌套或者继承

	interface A {
	  plus()
	}

	interface B extends A {
		minus()
	}

	class Test implement B {
		plus() {}
		minus() {}
	}

	Go语言中,是怎么做继承或者合并：嵌套方式

	type A interface {
		plus()
	}

	type B interface {
		A
		minus()
	}
*/

type User interface {
	login()
	register()
	error(err string)
}

type Admin interface {
	User
	admingLogin()
}

type SystemAdmin struct {
	username   string
	password   string
	passcode   string
	randomCode string
}

func (s *SystemAdmin) login() {
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

func (s *SystemAdmin) register() {
	// TODO implement me
	panic("implement me")
}

func (s *SystemAdmin) error(err string) {
	fmt.Println(err)
}

func (s *SystemAdmin) admingLogin() {
	// TODO implement me
	panic("implement me")
}

func (s *SystemAdmin) test() {
	s.username = "xulei1"
}

func main() {

	// &获取一个值的指针
	var admin Admin = &SystemAdmin{
		username:   "xulei11",
		password:   "123456",
		passcode:   "123456",
		randomCode: "dadasfadsfads",
	}

	admin.login()
}
