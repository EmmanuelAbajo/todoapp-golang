package dto

// TodoResponseDTO object
type TodoResponseDTO struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Content string `json:"content"`
}