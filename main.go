package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"todoapp/db"
	"todoapp/handlers"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	sqlHost       = "postgres"
	sqlConnString = "host=localhost user=postgres password=postgres dbname=todo port=5432 sslmode=disable"
)

func main() {
	conn, e := sql.Open(sqlHost, sqlConnString)
	if e != nil {
		panic(e)
	}
	defer conn.Close()

	repo := db.NewTodoRepo(conn)
	handler := handlers.NewTodoHandler(repo)

	router := mux.NewRouter()

	router.HandleFunc("/todos", handler.GetAllTodos)
	router.HandleFunc("/todos/{id}", handler.GetTodoById)
	fmt.Println("server is starting on 8000 \n ")
	if e = http.ListenAndServe(":8000", router); e != nil {
		panic(e)
	}

}
