package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var todoChan = make(chan []TodoItem, 1)

type TodoItem struct {
	id        int
	content   string
	completed bool
}

var todos = []TodoItem{}

func main() {

	// wg.Add(3)
	AddTodo(TodoItem{
		id:        1,
		content:   "text1",
		completed: false,
	})

	AddTodo(TodoItem{
		id:        2,
		content:   "text2",
		completed: false,
	})

	// go RemoveTodo(1)
	ToggleTodo(1)
	// wg.Wait()
	fmt.Println(34, todos)
}

func AddTodo(todoItem TodoItem) {
	// defer wg.Done()
	// 将 todoItem 放入 todoChan
	todoChan <- todos

	channelTodo, ok := <-todoChan
	if ok {
		todos = append(channelTodo, todoItem)
	}
}

func RemoveTodo(id int) {
	// defer wg.Done()
	defer close(todoChan)

	todoChan <- todos

	var delId int
	for i := 0; i < len(todos); i++ {

		if todos[i].id == id {
			delId = i
		}
	}
	fmt.Println("delete-id", delId)

	channelTodo, ok := <-todoChan
	if ok {
		if delId == 0 {
			todos = append(channelTodo[1:])
		} else if delId == len(channelTodo) {
			todos = append(channelTodo[:len(channelTodo)])
		} else {
			todos = append(channelTodo[0:delId], channelTodo[delId+1])
		}
	}
}

func ToggleTodo(id int) {
	// defer wg.Done()
	defer close(todoChan)

	todoChan <- todos

	for i := 0; i < len(todos); i++ {
		if todos[i].id == id {
			todos[i].completed = !todos[i].completed
		}
	}
}

func GetTodo(id int) {
	defer close(todoChan)
}
