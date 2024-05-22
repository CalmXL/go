package main

import (
	"fmt"
	"go-basics/8_concurrent_programming/todolist/todolist"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 主 Goroutine
func main() {
	wg.Add(1)
	go initialize()
	wg.Wait()
}

func initialize() {
	go showMenu()
	go todoListener()
}

/*
增加项目 1
删出项目 2
操作完成 3
查询所有项目 4
退出 5
*/
var todoList todolist.TodoInterface = &todolist.TodoList{
	List:     make([]todolist.Todo, 0),
	TodoChan: make(chan todolist.Action),
}

func showMenu() {
	num := 0
	fmt.Println(`
		增加项目 1
		删出项目 2
		操作完成 3
		查询所有项目 4
		退出 5
	`)
	fmt.Println("请输入操作类型编号")

	_, err := fmt.Scan(&num)
	if err != nil {
		return
	}

	switch num {
	case 1:
		// AddTodo => 输入 content
		todolist.AddTodoScan(todoList.Add)
	case 2:
		// RemoveTodo => 输入 Id
		todolist.RemoveTodoScan(todoList.Remove)
	case 3:
		// ToggleTodo => 输入 Id
		todolist.ToggleTodoScan(todoList.Toggle)
	case 4:
		// GetTodo
		todoList.Get()
	case 5:
		// Close
		todoList.Close()
	default:
	}
}

func todoListener() {
loop:
	for {
		select {
		case res, ok := <-todoList.GetChan():
			if !ok {
				break loop
			}
			switch res.Type {
			case todolist.TodoAdd:
				fmt.Println("您添加了如下的信息: ")
				todolist.PrintTodo(res.PayLoad.(todolist.Todo))
				go initialize()
			case todolist.TodoRemove:
				fmt.Println("请您输入你要移除的ID: ")
				todolist.PrintTodo(res.PayLoad.(todolist.Todo))
				go initialize()
			case todolist.TodoGet:
				fmt.Println("项目列表如下：")
				todolist.PrintTodoList(res.PayLoad.([]todolist.Todo))
				go initialize()
			case todolist.TodoStop:
				fmt.Println("您已退出!!")
				wg.Done()
			case todolist.TodoToggle:
				fmt.Println("请您输入你要操作的ID: ")
				todolist.PrintTodo(res.PayLoad.(todolist.Todo))
				go initialize()
			}
		default:
			time.Sleep(time.Second)
		}
	}
}
