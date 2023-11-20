package todolist

import "fmt"

type Todo struct {
	Id        int64
	Content   string
	Completed bool
}

func (td *Todo) ToString() {
	fmt.Printf(`
		ID: %d
		Content: %s
		completed: %t
	`, td.Id, td.Content, td.Completed)
}

func (td *Todo) toggle() {
	td.Completed = !td.Completed
}
