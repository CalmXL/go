package main

import "fmt"

func test1(a int) {
	a = 100
	fmt.Println("test1: ", a)
}

func test2(s []int) {
	s[0] = 100
	fmt.Println("test2: ", s)
	fmt.Printf("%p\r\n", s)
}

func test3(m map[int]int) {
	m[1] = 100
	fmt.Println("test3: ", m)
	fmt.Printf("%p\r\n", m)
}

func test4(fn func()) {
	fmt.Printf("%p\r\n", fn)
}

type Todo struct {
	id        int
	content   string
	completed bool
}

// 给结构体绑定函数
// 必须有一个函数的接收器(接受一个结构体)

func (todo *Todo) setId(id int) int {
	fmt.Printf("%p\r\n", todo)
	todo.id = id
	return id
}

func (todo *Todo) setContent(content string) string {
	fmt.Printf("%p\r\n", todo)
	todo.content = content
	return content
}

func (todo *Todo) setCompleted(completed bool) bool {
	fmt.Printf("%p\r\n", todo)
	todo.completed = completed
	return completed
}

type Todo2 struct {
	id        int
	content   string
	completed bool
}

func (todo Todo2) setId(id int) int {
	fmt.Printf("%p\r\n", &todo)
	todo.id = id
	return id
}

func (todo Todo2) setContent(content string) string {
	fmt.Printf("%p\r\n", &todo)
	todo.content = content
	return content
}

func (todo Todo2) setCompleted(completed bool) bool {
	fmt.Printf("%p\r\n", &todo)
	todo.completed = completed
	return completed
}

func main() {
	// a := 1
	// test1(a)
	// // 1. 基本类型作为参数传递 => 值传递
	// fmt.Println("main: ", a)
	
	// --------------------------
	// slice、map、channel 都不是引用传递，都是值传递
	
	// s := []int{1, 2, 3}
	// test2(s)
	// fmt.Println("main: ", s)
	// fmt.Printf("%p\r\n", s)
	
	// m := map[int]int{
	// 	1: 1,
	// 	2: 2,
	// 	3: 3,
	// }
	//
	// test3(m)
	// fmt.Println("main: ", m)
	// fmt.Printf("%p\r\n", m)
	
	// fn := func() {}
	// test4(fn)
	//
	// fmt.Printf("%p\r\n", fn)
	
	// todo := Todo{
	// 	id:        1,
	// 	content:   "todo1",
	// 	completed: false,
	// }
	//
	// todo.setId(2)
	// todo.setContent("change Todo1")
	// todo.setCompleted(!todo.completed)
	//
	// fmt.Println(todo)
	// fmt.Printf("%p\r\n", &todo)
	
	todo2 := Todo2{
		id:        2,
		content:   "todo2",
		completed: true,
	}
	
	todo2.setId(222)
	todo2.setContent("todo2 change")
	todo2.setCompleted(!todo2.completed)
	
	fmt.Println(todo2)
	fmt.Printf("%p\r\n", &todo2)
}
