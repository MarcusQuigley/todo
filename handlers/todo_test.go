package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"todoapp/models"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetAllTodos(t *testing.T) {
	want := `[{"id":1,"description":"first todo","is_completed":true},{"id":2,"description":"second todo","is_completed":false}]
`
	todos := []models.Todo{
		models.NewTodo(1, "first todo", true),
		models.NewTodo(2, "second todo", false),
	}
	th := NewTodoHandler(todos)
	r := httptest.NewRequest(http.MethodGet, "/todos", nil)
	w := httptest.NewRecorder()
	th.GetAllTodos(w, r)
	got, e := io.ReadAll(w.Result().Body)
	if e != nil {
		t.Fatalf("unexpected while reading body: %v", e)
	}
	w.Result().Body.Close()
	assert.Equal(t, want, string(got))
}

func TestGetTodoById(t *testing.T) {
	want := `{"id":1,"description":"first todo","is_completed":true}
`

	todos := []models.Todo{
		models.NewTodo(1, "first todo", true),
	}
	th := NewTodoHandler(todos)

	r := httptest.NewRequest(http.MethodGet, "/todos/1", nil)
	w := httptest.NewRecorder()
	mux.SetURLVars(r, map[string]string{"id": "1"})
	th.GetTodoById(w, r)
	got, e := io.ReadAll(w.Result().Body)
	if e != nil {
		panic(e)
	}
	assert.Equal(t, want, string(got))
}
