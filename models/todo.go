package models

type Todo struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`
}

func NewTodo(id int, description string, completed bool) Todo {
	return Todo{
		Id:          id,
		Description: description,
		IsCompleted: completed,
	}
}
