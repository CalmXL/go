package main

type TodoList[T any] interface {
	addTodo(todo T) bool
	removeTodo(id int) T
	toggleTodo(id int) bool
	setContent(id int, content string) bool
	getTodo(id int) T
	getList() []T
}
