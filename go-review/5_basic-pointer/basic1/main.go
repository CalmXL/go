package main

import "fmt"

// func test(a *int) {
// 	b := 100
// 	fmt.Printf("%p\r\n", a) // => 0x1400000e0b0
// 	a = &b
// 	fmt.Printf("%p\r\n", a) // => 0x1400000e0b8
// }

type Student struct {
	name  string
	age   int
	score float32
}

func (stu *Student) setScore(score float32) float32 {
	fmt.Printf("%p\r\n", stu) // => 0x14000090020
	stu.score = score
	return stu.score
}

func main() {
	// a := 1
	// test(&a)
	// fmt.Println(a) // => 1
	
	// 动态分配内存
	// new(类型) => 分配一个内存空间 => 内存空间的指针
	// a := new(int)
	// fmt.Println(a)
	//
	// *a = 100
	// fmt.Println(*a) // => 100 value variable
	
	// ------------------
	stu := Student{
		name:  "张三",
		age:   18,
		score: 81.5,
	}
	
	fmt.Printf("%p\r\n", &stu) // => 0x14000090020
	stu.setScore(99.5)
	fmt.Println(stu)
}
