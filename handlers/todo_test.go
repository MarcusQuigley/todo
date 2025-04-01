package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllTodos(t *testing.T) {
	want := `[{"id":1,"description":"first todo","is_completed":true},{"id":2,"description":"second todo","is_completed":false}]
`
	r := httptest.NewRequest(http.MethodGet, "/todos", nil)
	w := httptest.NewRecorder()
	GetAllTodos(w, r)
	got, e := io.ReadAll(w.Result().Body)
	if e != nil {
		t.Fatalf("unexpected while reading body: %v", e)
	}
	w.Result().Body.Close()
	assert.Equal(t, want, string(got))
}

func TestGetTodoById(t *testing.T) {
}
