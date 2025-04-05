package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todoapp/db"
)

type TodoHandler struct {
	rs db.IRepoService
}

func NewTodoHandler(service db.IRepoService) *TodoHandler {
	return &TodoHandler{service}
}

func (handler *TodoHandler) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos, e := handler.rs.GetAll()
	if e != nil {
		fmt.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	e = json.NewEncoder(w).Encode(todos)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *TodoHandler) GetTodoById(w http.ResponseWriter, r *http.Request) {
	todo, e := handler.rs.GetById(1)
	if e != nil {
		fmt.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	e = json.NewEncoder(w).Encode(todo)
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
