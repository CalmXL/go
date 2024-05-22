package main

import "fmt"

/*
  什么是接口(interface)？
		中国的插头 => 中国的电源（接口）
		英国的插头 => 英国的电源 (接口)
	接口就是一系列规范方法实现的标准

  // 设计方法由规范决定
	func add (a int, b int) int

  // 编码只需要按照规范来书写就行了
	func add () {}

	interface {
		sing (person, whichSong) voice
		run  (person, meters) calorie
	}

	func sing (p Person, whichSong Song) voice {}

	interface 也是自定义类型
	type Duck interface {

	}

	空接口表示任何类型都可以接收 => any
	接口还可以描述泛型的类型约束
*/

type MyInt int
type MyFloat float64
type PlusT interface {
	// ~ 是包含衍生类型(相关的自定义类型的含义)
	~int | string | float32 | ~float64 | uint8
}

// func plus[T int | string | float64 | float32 | uint8](a, b T) T {
// }

func plus[T PlusT](a, b T) T {
	return a + b
}

type Duck interface {
	walk()
	shout()
}

type Duck2 interface {
	walk()
}

type Duck3 interface {
	shout()
}

type Bird struct {
	legs int
}

type Dog struct {
	legs int
}

func (bd *Bird) walk() {
	fmt.Println("Bird is walking")
}

func (bd *Bird) shout() {
	fmt.Println("Bird is shouting")
}

func (dg *Dog) walk() {
	fmt.Println("Dog is walking")
}

func (dg *Dog) shout() {
	fmt.Println("Dog is shouting")
}

func (dg *Dog) test() {
	fmt.Println("Dog is testing")
}

func doSth(animal Duck) {
	animal.walk()
	animal.shout()
}

func main() {
	// any => interface{} => 可以接受任意类型
	// studentMap := map[string]any{
	// 	"name": "张三",
	// 	"age":  18,
	// }

	// ------------------------------------------
	// var a MyInt = 1
	// var b MyInt = 2
	// var c = 3
	// var d = 4
	//
	// plus(a, b)
	// plus(c, d)

	// var a MyFloat = 1
	// var b MyFloat = 2
	//
	// plus(a, b)

	// --------------------------------------

	// doSth(&Bird{
	// 	legs: 2,
	// })
	//
	// doSth(&Dog{
	// 	legs: 4,
	// })

	var bird Duck2 = &Bird{legs: 2}
	var dog Duck3 = &Dog{legs: 4}

	bird.walk()
	dog.shout()
}
