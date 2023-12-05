package todolist

import "fmt"

func PrintTodo(td Todo) {
	fmt.Printf(`
		ID: %d
		Content: %s
		completed: %t
	`, td.Id, td.Content, td.Completed)
}

func PrintTodoList(todoList []Todo) {
	for _, el := range todoList {
		PrintTodo(el)
	}
}
