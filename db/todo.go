package db

import (
	"database/sql"

	"todoapp/models"
)

func GetAll(db *sql.DB) ([]models.Todo, error) {
	rows, _ := db.Query("SELECT id, description, is_completed FROM todos")
	defer rows.Close()
	return []models.Todo{
		models.NewTodo(1, "first todo", false),
		models.NewTodo(2, "second todo", true),
	}, nil
	//return nil, nil
}
