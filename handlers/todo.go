package handlers

import (
	"encoding/json"
	"net/http"
	"todoapp/models"
)

type TodoHandler struct {
	todos []models.Todo
}

func NewTodoHandler(todos []models.Todo) *TodoHandler {
	return &TodoHandler{todos}
}

func (td *TodoHandler) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	e := json.NewEncoder(w).Encode(td.todos)
	if e != nil {
		panic(e)
	}
}

func (td *TodoHandler) GetTodoById(w http.ResponseWriter, r *http.Request) {
	e := json.NewEncoder(w).Encode(td.todos[0])
	if e != nil {
		panic(e)
	}
}
