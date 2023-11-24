package main

type SchedList struct {
	list []Schedule
}

func (sl *SchedList) addTodo(todo Schedule) bool {
	for _, el := range sl.list {
		if el.id == todo.id {
			return false
		}
	}

	sl.list = append(sl.list, todo)
	return true
}

func (sl *SchedList) removeTodo(id int) (element Schedule) {
	for i, el := range sl.list {
		if el.id == id {
			element = el
			sl.list = append(sl.list[:i], sl.list[i+1:]...)
		}
	}
	return element
}

func (sl *SchedList) toggleTodo(id int) bool {
	for _, el := range sl.list {
		if el.id == id {
			el.completed = !el.completed
			return true
		}
	}
	return false
}

func (sl *SchedList) setContent(id int, content string) bool {
	for _, el := range sl.list {
		if el.id == id {
			el.content = content
			return true
		}
	}
	return false
}
func (sl *SchedList) getTodo(id int) (element Schedule) {
	for _, el := range sl.list {
		if el.id == id {
			element = el
		}
	}
	return element
}
func (sl *SchedList) getList() []Schedule {
	return sl.list
}
