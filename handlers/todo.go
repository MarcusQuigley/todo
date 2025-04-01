package handlers

import (
    "net/http"
    "encoding/json"
 )

type Todo struct{
    Id int `json:"id"`
    Description string `json:"description"`
    IsCompleted bool `json:"is_completed"`
}

func GetAllTodos(w http.ResponseWriter, r *http.Request){
    todos:= [2]Todo{
        // NewTodo(1, "first todo", true),
        // NewTodo(2, "second todo", false),
        {1, "first todo", true},
        {2, "second todo", false},
    }

    e:=json.NewEncoder(w).Encode(todos)
    if e!=nil{
        panic(e)
    }

    //w.Write()
}