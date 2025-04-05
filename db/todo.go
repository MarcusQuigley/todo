package db

import (
	"database/sql"
	"todoapp/models"
)

const (
	SqlSelectAllTodos = "SELECT id, description, is_completed FROM todos"
	SqlSelectATodo    = "SELECT id, description, is_completed FROM todos WHERE id = $1"
)

type IRepoService interface {
	GetAll() ([]models.Todo, error)
	GetById(id int) (models.Todo, error)
}

func NewTodoRepo(db *sql.DB) *TodoRepo {
	return &TodoRepo{db}
}

type TodoRepo struct {
	db *sql.DB
}

func (tr *TodoRepo) GetAll() ([]models.Todo, error) {
	var todos []models.Todo
	rows, e := tr.db.Query(SqlSelectAllTodos)
	if e != nil {
		return todos, e
	}
	defer rows.Close()

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

func (tr *TodoRepo) GetById(id int) (models.Todo, error) {
	row := tr.db.QueryRow(SqlSelectATodo, id)
	var todo models.Todo

	e := row.Scan(&todo.Id, &todo.Description, &todo.IsCompleted)

	return todo, e

}
