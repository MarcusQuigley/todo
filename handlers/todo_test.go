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

type repoMock struct {
	fnGetAll  func() ([]models.Todo, error)
	fnGetById func(id int) (models.Todo, error)
}

func (rm *repoMock) GetAll() ([]models.Todo, error) {
	return rm.fnGetAll()
}

func (rm *repoMock) GetById(id int) (models.Todo, error) {
	return rm.fnGetById(id)
}

func TestGetAllTodos(t *testing.T) {
	want := `[{"id":1,"description":"first todo","is_completed":true},{"id":2,"description":"second todo","is_completed":false}]
`
	mock := &repoMock{
		fnGetAll: func() ([]models.Todo, error) {
			return []models.Todo{
				models.NewTodo(1, "first todo", true),
				models.NewTodo(2, "second todo", false),
			}, nil
		},
	}

	th := NewTodoHandler(mock)
	r := httptest.NewRequest(http.MethodGet, "/todos", nil)
	w := httptest.NewRecorder()
	th.GetAllTodos(w, r)
	got, e := io.ReadAll(w.Result().Body)
	if e != nil {
		t.Fatalf("unexpected while reading body: %v", e)
	}
	w.Result().Body.Close()
	assert.Equal(t, string(got), want)
}

func TestGetTodoById(t *testing.T) {
	want := `{"id":1,"description":"first todo","is_completed":true}
`
	mock := &repoMock{
		fnGetById: func(id int) (models.Todo, error) {
			return models.NewTodo(id, "first todo", true), nil
		},
	}

	th := NewTodoHandler(mock)

	r := httptest.NewRequest(http.MethodGet, "/todos/1", nil)
	w := httptest.NewRecorder()

	r = mux.SetURLVars(r, map[string]string{"id": "1"})
	th.GetTodoById(w, r)

	got, e := io.ReadAll(w.Result().Body)
	if e != nil {
		panic(e)
	}
	w.Result().Body.Close()
	assert.Equal(t, string(got), want)
}
