package db

import (
	"database/sql"

	"todoapp/models"
)

const (
	SqlSelectAllTodos = "SELECT id, description, is_completed FROM todos"
	SqlSelectATodo    = "SELECT id, description, is_completed FROM todos where id = $1"
)

func GetAll(db *sql.DB) ([]models.Todo, error) {
	rows, _ := db.Query(SqlSelectAllTodos)
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
func GetByID(id int, db *sql.DB) (models.Todo, error) {
	row := db.QueryRow(SqlSelectATodo, id)
	var todo models.Todo
	row.Scan(&todo.Id, &todo.Description, &todo.IsCompleted)
	e := row.Err()
	return todo, e
}
