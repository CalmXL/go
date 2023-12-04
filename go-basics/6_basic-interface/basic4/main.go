package main

import "fmt"

func main() {
	var schedList TodoList[Schedule] = &SchedList{
		list: []*Schedule{},
	}

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

	sl := schedList.getList()
	fmt.Println(sl)

	// res1 := schedList.removeTodo(1)
	// fmt.Println(res1)
	// fmt.Println(schedList.getList())
	//
	// res2 := schedList.toggleTodo(2)
	// fmt.Println(res2)
	// fmt.Println(schedList.getList())
	//
	// res3 := schedList.setContent(3, "new todo3")
	// fmt.Println(res3, schedList.getList())
}
