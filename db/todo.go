package db

import (
	"database/sql"

	"todoapp/models"
)

const (
	SqlSelectAllTodos = "SELECT id, description, is_completed FROM todos"
	SqlSelectATodo    = "SELECT id, description, is_completed FROM todos where id = $1"
)

type TodoHandler struct {
	db *sql.DB
}

func NewTodoHandler(db *sql.DB) TodoHandler {
	return TodoHandler{db}
}

func (td *TodoHandler) GetAll() ([]models.Todo, error) {
	rows, _ := td.db.Query(SqlSelectAllTodos)
	defer rows.Close()
	var todos []models.Todo
	for rows.Next() {
		var td models.Todo
		e := rows.Scan(&td.Id, &td.Description, &td.IsCompleted)
		if e != nil {
			continue
		}
		todos = append(todos, td)
	}
	return todos, nil
}
func (td *TodoHandler) GetByID(id int) (models.Todo, error) {
	row := td.db.QueryRow(SqlSelectATodo, id)
	var todo models.Todo
	row.Scan(&todo.Id, &todo.Description, &todo.IsCompleted)
	e := row.Err()
	return todo, e
}
