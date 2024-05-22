package main

import "fmt"

// 一个函数的参数，可以理解为函数的局部变量
func exchange(x *int, y *int) {
	temp := x
	x = y
	y = temp
	
	fmt.Println(x, y)
	fmt.Println(*x, *y) // 2 1
}

func main() {
	
	// x := 1
	// y := 2
	//
	// exchange(&x, &y)
	// fmt.Println(&x, &y)
	//
	// var a *int
	// b := 100
	// a = &b
	//
	// fmt.Println(a)
	// fmt.Printf("%p\r\n", a)
	
	// a := 1
	// var b *int
	// b = &a
	//
	// fmt.Println(a, b)
	
	// a := new(int) // 返回一个指针的类型，不是野指针
	// fmt.Println(a)
	// fmt.Println(*a)
	// fmt.Printf("%T %p", a, a)
	
	// ============================
	// type StudentMap struct {
	// 	name string
	// 	age  int
	// }
	//
	// stuPointer := &StudentMap{
	// 	name: "张三",
	// 	age:  18,
	// }
	//
	// setName := func(stuPointer *StudentMap, name string) {
	// 	stuPointer.name = name
	// }
	//
	// setName(stuPointer, "李四")
	//
	// fmt.Println(*stuPointer)
	
	// ===============================
	type StuMap map[string]interface{}
	
	studentMap := StuMap{
		"name": "张三",
		"age":  18,
	}
	
	setName := func(stu StuMap, name string) {
		stu["name"] = name
		fmt.Println(&stu)
		fmt.Printf("%p\r\n", stu)
	}
	
	setName(studentMap, "流水")
	fmt.Println(studentMap)
}
