package main

import "fmt"

/*
type 关键字： 主要是用于对类型进行定义的工具
 1. 给一个类型起一个别名
 2. 创建一个新的类型
*/
func main() {
	/**
	  账户(string)、户名(string)、有效性(bool)、余额(float32)、总额(float32)
	*/
	// type 的名称首字母必须大写
	// type TypeAccountNumber = string
	// 类型别名的作用：主要是让类型具备语义化
	// var accNum TypeAccountNumber = "1234 5678"
	// accountName := "小野森森"
	// accountValid := true
	// accountBalance := 10000

	// ------------------------------------------------
	// 类型别名的作用: 主要是给变量指定类型的时候可以简短一些
	// type TypeCallback = func(a int, b int, sign string, res string) string
	// type TypeStudentCollection = []map[string]interface{}
	//
	// compute := func(a int, b int, method string, cb TypeCallback) {
	//
	// }
	//
	// test := func(studentCollection TypeStudentCollection) {}

	// ------------------------------------------------
	// type TypeTodo = map[string]interface{}
	// type TypeTodoList = []TypeTodo
	// var todoList TypeTodoList
	//
	// setTodo := func(id int, content string, completed bool) TypeTodo {
	// 	todoMap := TypeTodo{
	// 		"id":        id,
	// 		"content":   content,
	// 		"completed": completed,
	// 	}
	//
	// 	return todoMap
	// }
	//
	// appendTodo := func(todoList TypeTodoList, todos ...TypeTodo) TypeTodoList {
	// 	return append(todoList, todos...)
	// }
	//
	// todo1 := setTodo(1, "todo1", false)
	// todo2 := setTodo(2, "todo2", true)
	// todo3 := setTodo(3, "todo3", false)
	//
	// todoList = appendTodo(todoList, todo1, todo2, todo3)
	//
	// fmt.Println(todoList)

	// -------------------------------------------------------
	/**
	  var  a number // 自定义变量
		type a number // 自定义类型
		类型a
			1. 类型 a 具备 number 的所有特性
			2. 类型 a 是 Go 语言中被创建的真实类型
	*/
	type TypeTodo map[string]interface{}
	type TypeTodoList []TypeTodo
	var todoList TypeTodoList

	setTodo := func(id int, content string, completed bool) TypeTodo {
		todoMap := TypeTodo{
			"id":        id,
			"content":   content,
			"completed": completed,
		}

		return todoMap
	}

	appendTodo := func(todoList TypeTodoList, todos ...TypeTodo) TypeTodoList {
		return append(todoList, todos...)
	}

	todo1 := setTodo(1, "todo1", false)
	todo2 := setTodo(2, "todo2", true)
	todo3 := setTodo(3, "todo3", false)

	todoList = appendTodo(todoList, todo1, todo2, todo3)

	fmt.Println(todoList)
	fmt.Printf("%T", todoList) // => main.TypeTodoList

	/**
	  类型别名与自定义类型的区别：
			1. 类型别名是为了对比较长的或者需要语义化的类型进行取一个名字（短的）
			2. 自定义类型是创建一种新的类型，有原来的类型的特征，但又是一个独立的类型
			3. 类型别名的类型还是等号右边的类型
			4. 自定义类型的类型就是定义的类型
			5. 如果仅仅为了语义化类型定义或者短化类型定义，那么就选类型别名
			6. 如果一种类型的定义是复合型的、多元素的、复杂的，那么就选自定义类型
	*/
}
