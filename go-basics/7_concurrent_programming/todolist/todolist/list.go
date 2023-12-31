package todolist

import "time"

type OpType int

const (
	TodoGet OpType = iota
	TodoAdd
	TodoRemove
	TodoToggle
	TodoStop
)

type Action struct {
	Type    OpType
	PayLoad any
}

type TodoInterface interface {
	Get()                 // 获取列表
	Add(content string)   // 增加项
	Remove(id int64)      // 删除项
	Toggle(id int64)      // 完成项
	GetChan() chan Action // 获取 channel
	GetTodoList() []Todo
	Close() // 关闭 channel
}

type TodoList struct {
	List     []Todo
	TodoChan chan Action
}

func (ls *TodoList) Get() {
	ls.TodoChan <- Action{
		Type:    TodoGet,
		PayLoad: ls.List,
	}
}

func (ls *TodoList) Add(content string) {
	todo := Todo{
		Id:        time.Now().Unix(),
		Content:   content,
		Completed: false,
	}

	ls.List = append(ls.List, todo)
	ls.TodoChan <- Action{
		Type:    TodoAdd,
		PayLoad: todo,
	}
}

func (ls *TodoList) Remove(id int64) {
	for i, todo := range ls.List {
		if todo.Id == id {
			ls.List = append(ls.List[:i], ls.List[(i+1):]...)
			ls.TodoChan <- Action{
				Type:    TodoRemove,
				PayLoad: todo,
			}
		}
	}
}

func (ls *TodoList) Toggle(id int64) {
	// for range => todo 是原本 List 里面的 todo 的副本
	for i, todo := range ls.List {
		if todo.Id == id {
			// ls.list[i] => go 语言赋值 值传递，并不是那个引用
			_todo := &ls.List[i]
			_todo.toggle()
			ls.TodoChan <- Action{
				Type:    TodoToggle,
				PayLoad: *_todo,
			}
		}
	}
}

func (ls *TodoList) GetChan() chan Action {
	return ls.TodoChan
}

func (ls *TodoList) GetTodoList() []Todo {
	return ls.List
}

func (ls *TodoList) Close() {
	ls.TodoChan <- Action{
		Type:    TodoStop,
		PayLoad: Todo{},
	}
	close(ls.TodoChan)
}
