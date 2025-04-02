package db

import (
	"database/sql"

	"todoapp/models"
)

const SqlSelectAllTodos = "SELECT id, description, is_completed FROM todos"

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
