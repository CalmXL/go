package main

import "fmt"

type ScheduleList struct {
	list []*Schedule
}

func (sl *ScheduleList) addTodo(todo *Schedule) bool {
	
	for _, el := range sl.list {
		if el.id == todo.id {
			return false
		}
	}
	
	sl.list = append(sl.list, todo)
	return true
}

func (sl *ScheduleList) removeTodo(id int) (element Schedule) {
	for index, el := range sl.list {
		if el.id == id {
			element = *el
			sl.list = append(sl.list[:index], sl.list[index+1:]...)
		}
	}
	
	return element
}

func (sl *ScheduleList) toggleTodo(id int) bool {
	for _, el := range sl.list {
		if el.id == id {
			// sl.list[index].completed = !sl.list[index].completed
			el.completed = !el.completed
			return true
		}
	}
	
	return false
}

func (sl *ScheduleList) setContent(id int, content string) bool {
	for _, el := range sl.list {
		if el.id == id {
			el.content = content
			fmt.Println("el:", el)
			fmt.Println("sl:", sl)
			return true
		}
	}
	return false
}

func (sl *ScheduleList) getTodo(id int) (element Schedule) {
	for _, el := range sl.list {
		if el.id == id {
			element = *el
		}
	}
	return element
}

func (sl *ScheduleList) getList() []Schedule {
	newSlice := make([]Schedule, len(sl.list))
	
	for index, el := range sl.list {
		newSlice[index] = *el
	}
	
	return newSlice
}
