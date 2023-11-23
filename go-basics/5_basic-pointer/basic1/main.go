package main

import "fmt"

/*
  1.什么是指针？
		指针在 Go 语言中，是一种存放内存空间地址的类型（数据类型）
		指针总会指向一块特定的空间（地址对应的空间）
		在 Go 语言中，指针可以访问其对应内部的属性
		指针在 Go 语言中不能直接参与运算 (unsafe.Pointer)
*/
/*
	2. 指针的基本使用：
		指针的类型: *普通数据类型   *int  *bool
		指针值: &普通数据变量名     &a    &b
		指针获取值: *指针变量名     *a    *b

		普通值转指针: &
		指针转普通值: *
		类型标识:     *

	3. 指针的作用？
		1. 动态分配内存（非常少地使用情况）
		2. 进行指针的传递
		3. 修改函数外部变量值或者其内部的变量值
*/

// int 类型数据的指针类型（一个空间的地址）
// func test(a *int) {
// 	b := 100
// 	fmt.Printf("%p\r\n", a)
// 	a = &b
// 	fmt.Printf("%p\r\n", a)
// }

// func test(a *int) {
// 	*a = 100
// }

type Student struct {
	name  string
	age   int
	score float32
}

func (stu *Student) setScore(score float32) float32 {
	fmt.Printf("%p\r\n", stu)
	stu.score = score
	return stu.score
}

func main() {
	// a := 1 // 有地址
	// test(&a)
	// fmt.Println(a)

	// ---------------------------------------------------------
	// 	a := 1
	// 	test(&a)
	// 	fmt.Println(a)

	// ---------------------------------------------------------
	// 动态分配内存空间
	// new(类型) => 分配出一个内存空间 => 内存空间的指针
	// a := new(int) // 返回 pointer
	// *a = 100      // value variable
	//
	// fmt.Println(*a)
	// b := 200
	// a = &b
	// fmt.Println(*a)

	// ---------------------------------------------------------
	// 指针传递
	stu := Student{
		name:  "张三",
		age:   18,
		score: 81.5,
	}
	fmt.Printf("%p\r\n", &stu)

	stu.setScore(90.5)
	fmt.Println(stu)
}
