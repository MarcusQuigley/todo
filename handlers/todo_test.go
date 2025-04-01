package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	//"github.com/stretchr/testify/assert"
)

func TestGetAllTodos(t *testing.T) {
	want := `[{"id":1,"description":"first todo","is_completed":true}]`
	r := httptest.NewRequest(http.MethodGet, "/todos", nil)
	w := httptest.NewRecorder()
	GetAllTodos(w, r)
	got, e := io.ReadAll(w.Result().Body)
	if e != nil {
		t.Fatalf("unexpected while reading body: %v", e)
	}
	w.Result().Body.Close()
	//assert.Equal(t, string(got), want)
	if string(got) != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGetTodoById(t *testing.T) {
	// for Task 2, write here the solution
}
