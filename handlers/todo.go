package handlers

import (
	"encoding/json"
	"net/http"
)

type Todo struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`
}

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos := []Todo{
		{1, "first todo", true},
		{2, "second todo", false},
	}

	e := json.NewEncoder(w).Encode(todos)
	if e != nil {
		panic(e)
	}
}
