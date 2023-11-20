package main

import "fmt"

/*
*

	  Go 语言中并没有面相对象的思想设计
	  class Object => 对象性质的设计是没有的
		结构体可以实现面相对象的一部分特征

	  结构体描述的是一组数据的结构，是一种 Go 语言层面上的自定义类型
		可以认为是多个字段的结构
		作用: 封装多种类型的数据
		类似：type AccountNumber string
		type StructName struct {
			// 结构体字段
			a int
			b int
			c bool
		}
*/
func main() {
	// todo1 := []string{"1", "todo1", "false"}
	// todo2 := []string{"2", "todo2", "true"}
	// todo3 := []string{"3", "todo3", "false"}
	//
	// todoList := [][]string{
	// 	todo1, todo2, todo3,
	// }
	//
	// fmt.Println(todoList)

	// -------------------------------------------
	// todo1 := map[string]interface{}{
	// 	"id":        1,
	// 	"content":   "todo1",
	// 	"completed": false,
	// }
	// todo2 := map[string]interface{}{
	// 	"id":        2,
	// 	"content":   "todo2",
	// 	"completed": true,
	// }
	// todo3 := map[string]interface{}{
	// 	"id":        3,
	// 	"content":   "todo3",
	// 	"completed": false,
	// }
	//
	// todoList := []map[string]interface{}{
	// 	todo1, todo2, todo3,
	// }
	//
	// fmt.Println(todoList)

	// -----------------------------------------
	// type TodoStruct struct {
	// 	id        int
	// 	content   string
	// 	completed bool
	// }
	//
	// todo1 := TodoStruct{
	// 	1,
	// 	"todo1",
	// 	false,
	// }
	// todo2 := TodoStruct{
	// 	2,
	// 	"todo2",
	// 	true,
	// }
	// 最推荐的
	/**
	  1. 字段名清晰，易读
		2. 可以不完整的进行赋值
	*/
	// todo3 := TodoStruct{
	// 	id:        3,
	// 	content:   "todo1",
	// 	completed: false,
	// }
	//
	// // fmt.Println(todo3)
	//
	// todo2.content = "This is todo2"
	// todo2.completed = false
	//
	// todoList := []TodoStruct{
	// 	todo1, todo2, todo3,
	// }

	// fmt.Println(todoList)

	// ---------------------------------------
	// type Student struct {
	// 	name   string
	// 	age    int
	// 	course string
	// }
	//
	// students := []Student{
	// 	// 初始化结构体
	// 	{
	// 		name: "张三",
	// 		age:  18,
	// 	},
	// 	{
	// 		name:   "李四",
	// 		course: "Java WEB",
	// 	},
	// 	{
	// 		name: "王五",
	// 	},
	// 	{},
	// }
	//
	// fmt.Println(students)

	// -----------------------------------
	// type Todo struct {
	// 	id        int
	// 	content   string
	// 	completed bool
	// }

	// new 结构体声明初始化
	// todo := new(Todo)
	//
	// // 此时todo 是指针的类型
	// todo.id = 1
	// todo.content = "This is Todo"
	// todo.completed = false

	// ---------------------------------------------------

	/**
	{
		"stu1": {
			name,
			age,
			course
		}
	}
	*/
	type Student struct {
		name   string
		age    int
		course string
	}

	students := map[int]Student{
		1: {
			name:   "张三",
			age:    18,
			course: "JavaWEB",
		},
		2: {
			name:   "李四",
			age:    22,
			course: "Golang",
		},
		3: {
			name:   "王五",
			age:    26,
			course: "Rust",
		},
	}

	fmt.Println(students)

	for k, v := range students {
		fmt.Println(k, v)
	}

	fmt.Println(students[1].name)
}
