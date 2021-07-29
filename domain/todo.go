package domain

import "fmt"

// Todo Represents the todo model
type Todo struct {
	ID int `gorm:"primary_key" json:"id"`
	Name string `gorm:"type:varchar(100)" json:"name"`
	Content string `gorm:"type:varchar(100)" json:"content"`
}

func (todo *Todo) toString() string {
	return fmt.Sprintf("Todo: {id: %d, name: %s, content: %s }", todo.ID, todo.Name, todo.Content)
}
