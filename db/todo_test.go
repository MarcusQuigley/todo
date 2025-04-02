package db

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	db, mock, e := sqlmock.New()
	if e != nil {
		t.Fatalf("error %v not expected while opening a connection to a mock", e)
	}
	defer db.Close()
	td := NewTodoHandler(db)
	rows := sqlmock.NewRows([]string{"id", "description", "is_completed"}).AddRow(1, "first todo", false).AddRow(2, "second todo", true)
	mock.ExpectQuery("SELECT id, description, is_completed FROM todos").WillReturnRows(rows)
	todos, er := td.GetAll()
	if er != nil {
		t.Fatalf("error %v not expected while calling GetAll", er)
	}
	assert.Equal(t, 2, len(todos))
	if e = mock.ExpectationsWereMet(); e != nil {
		t.Errorf("there were unfulfilled expectations: %v", e)
	}
}

func TestGetByID(t *testing.T) {
	id := 1
	db, mock, e := sqlmock.New()
	if e != nil {
		t.Fatalf("error %v not expected while opening a connection to a mock", e)
	}
	defer db.Close()
	td := NewTodoHandler(db)
	row := sqlmock.NewRows([]string{"id", "description", "is_completed"}).AddRow(1, "first todo", true)

	mock.ExpectQuery("SELECT id, description, is_completed FROM todos where id = ?").WithArgs(id).WillReturnRows(row)
	todo, e := td.GetByID(id)
	if e != nil {
		t.Fatalf("error %v not expected while calling GetByID", e)
	}
	assert.NotNil(t, todo)
	if e = mock.ExpectationsWereMet(); e != nil {

	}
}
