package main

import "fmt"

func main() {
	// sl := ScheduleList{
	// 	list: []Schedule{},
	// }
	//
	// fmt.Println(sl.list)
	//
	// sl.addTodo(Schedule{
	// 	id:        1,
	// 	content:   "todo1",
	// 	completed: false,
	// })
	//
	// sl.addTodo(Schedule{
	// 	id:        2,
	// 	content:   "todo2",
	// 	completed: false,
	// })
	//
	// sl.addTodo(Schedule{
	// 	id:        3,
	// 	content:   "todo3",
	// 	completed: true,
	// })
	//
	// fmt.Println(sl.list)
	
	var schedList TodoList[Schedule] = &ScheduleList{
		list: []*Schedule{},
	}
	
	// struct 值传递，所以传递指针
	schedList.addTodo(&Schedule{
		id:        1,
		content:   "todo1",
		completed: false,
	})
	
	schedList.addTodo(&Schedule{
		id:        2,
		content:   "todo2",
		completed: true,
	})
	
	schedList.addTodo(&Schedule{
		id:        3,
		content:   "todo3",
		completed: false,
	})
	
	// fmt.Println(schedList.getList())
	//
	// schedList.removeTodo(1)
	//
	// fmt.Println(schedList.getList())
	//
	// schedList.toggleTodo(2)
	
	schedList.setContent(2, "this is todo2")
	fmt.Println(schedList.getList())
	
}
