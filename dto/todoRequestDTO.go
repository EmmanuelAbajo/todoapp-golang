package dto

// TodoRequestDTO object
// Fields must start with a upper case
type TodoRequestDTO struct {
	Name string `json:"name"`
	Content string `json:"content"`
}