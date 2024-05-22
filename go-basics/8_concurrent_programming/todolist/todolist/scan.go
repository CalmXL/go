package todolist

import "fmt"

func AddTodoScan(addTodo func(string)) {
	content := ""
	fmt.Println("请输入TODO内容: ")
	_, err := fmt.Scan(&content)
	if err != nil {
		return
	}

	go addTodo(content)
}

func RemoveTodoScan(removeTodo func(int64)) {
	var id int64 = 0
	fmt.Println("请输入要删除的项目Id: ")
	_, err := fmt.Scan(&id)
	if err != nil {
		return
	}

	go removeTodo(id)
}

func ToggleTodoScan(toggleTodo func(int64)) {
	var id int64 = 0
	fmt.Println("请输入要操作的项目Id: ")
	_, err := fmt.Scan(&id)
	if err != nil {
		return
	}

	go toggleTodo(id)
}
