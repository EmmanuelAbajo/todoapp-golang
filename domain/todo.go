package model

import "fmt"

// Todo Represents the todo model
type Todo struct {
	id int
	name string
	content string
}

func (todo *Todo) toString() string {
	return fmt.Sprintf("Todo: {id: %d, name: %s, content: %s }", todo.id, todo.name, todo.content)
}
