package main

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/todos", GetTodosHandler).Methods("GET")
	r.HandleFunc("/api/v1/todos", CreateTodoHandler).Methods("POST")
	return r
}
