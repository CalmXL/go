package main

import "fmt"

// 一个函数的参数，可以理解为函数的局部变量
// func exchange(x *int, y *int) {
// 	// temp := x
// 	// x = y
// 	// y = temp
//
// 	// x, y = y, x
// 	*x, *y = *y, *x
//
// 	fmt.Println(x, y)
// 	fmt.Println(*x, *y)
// }

func main() {
	// x := 1
	// y := 2
	//
	// exchange(&x, &y)
	// fmt.Println(&x, &y)
	// fmt.Println(x, y)

	// -----------------------------------------
	/*
		野指针: 这个指针不明确指向任何空间
	*/
	// 只是声明了一个指针变量，而并没有初始化
	// 但未初始化的指针是可以访问的
	// 访问的结果就是野指针
	// var a *int
	// b := 100
	// a = &b
	// fmt.Println(a)
	// // nil 0x0 这就是野指针的表现
	// fmt.Printf("%p", a)

	// -----------------------------------------
	// a := 1
	// // b := &a
	// var b *int
	// b = &a
	// fmt.Println(a, b)
	// ------------------------
	// 推荐使用
	// a := new(int) // 返回一个指针的类型， 不是野指针
	// fmt.Println(a)
	// fmt.Println(*a) // int 类型的默认值是0
	//
	// type Student struct {
	// 	name string
	// 	age  int
	// }
	//
	// stuPointer := &Student{
	// 	name: "张三",
	// 	age:  18,
	// }
	//
	// setName := func(stuPointer *Student, name string) {
	// 	stuPointer.name = name
	// }
	//
	// setName(stuPointer, "李四")
	// fmt.Println(*stuPointer)

	// ----------------------------------------------------------
	type StudentMap map[string]interface{}

	studentMap := StudentMap{
		"name": "张三",
		"age":  18,
	}

	setName := func(stu StudentMap, name string) {
		stu["name"] = name
		// // stu.name = name
		// fmt.Println(*stu)
		// fmt.Printf("%p\r\n", stu)

	}

	setName(studentMap, "李四")
	fmt.Println(studentMap)
}
