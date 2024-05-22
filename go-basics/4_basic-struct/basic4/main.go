package main

import "fmt"

// func test1(a int) {
//	a = 100
// }
//
// func test2(s []int) {
//	s[0] = 100
//	fmt.Println(s)

//	fmt.Printf("%p\r\n", s)
// }
//
// func test3(m map[int]int) {
//	m[1] = 100
//	fmt.Println(m)
//	fmt.Printf("%p\r\n", m)
// }
//
// func test4(f func()) {
//	fmt.Printf("%p\r\n", f)
// }

type Todo struct {
	id        int
	content   string
	completed bool
}

// 给结构体绑定函数
// 必须有一个函数的接收器 (接收一个结构体)
/**
  函数的 struct 接收器接收的是结构体的值，本质上是一种值传递
	结构体会复制一份传递给结构体的绑定函数
*/
/**
  值传递 => 传进来的是 todo 的副本，做更新的时候不会影响原数据，
						struct 传递是值传递，不是指针传递
	指针传递 => 传进来的 todo, 是 todo 值的指针，做更新的时候，会影响原数据

	所以如果给结构体绑定的函数内部仅仅是访问数据，就可以保持普通值的传递
	如果要在给结构体绑定的函数内部做字段值的操作，就一定要使用指针传递。

	在 Go 语言中，指针是可以直接访问值的，这是系统为了这个操作做的相应功能，
	其他静态语言是没有的
*/
func (todo *Todo) setId(id int) int {
	fmt.Printf("%p\r\n", todo)
	todo.id = id
	return id
}

func (todo *Todo) setContent(content string) string {
	todo.content = content
	return content
}

func (todo *Todo) setCompleted(completed bool) bool {
	todo.completed = completed
	return completed
}

func main() {
	// 基本类型作为参数传递 => 值传递
	/**
	  1. 复制一份数据
	  2. 将数据传递给函数
	*/
	// a := 1
	// test1(a)
	// fmt.Println(a)

	// -----------------------------
	// slice map channel 都不是引用传递， 都是值传递
	/**
	  在 Go 语言当中，虽然传递的时候没有显示的传递指针，
		但是在 Go 底层，会复制一份值的指针，传递给函数
	*/
	// s := []int{1, 2, 3}
	// test2(s)
	// fmt.Println(s)
	// fmt.Printf("%p\r\n", s)

	// m := map[int]int{
	//	1: 1,
	//	2: 2,
	// }
	//
	// test3(m)
	// fmt.Println(m)
	//
	// fmt.Printf("%p\r\n", m)

	// f := func() {}
	// test4(f)
	// fmt.Printf("%p\r\n", f)

	todo := Todo{
		id:        1,
		content:   "content1",
		completed: false,
	}

	fmt.Printf("%p\r\n", &todo)
	todo.setId(2)
	todo.setContent("this is content")
	todo.setCompleted(true)

	fmt.Println(todo)

}
