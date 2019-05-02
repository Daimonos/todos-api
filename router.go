package main

import (
	"github.com/gorilla/mux"
)

// NewRouter creates a new router for the application
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/todos", chainMiddleware(GetTodosHandler, checkAuthentication)).Methods("GET")
	r.HandleFunc("/api/v1/todos", chainMiddleware(CreateTodoHandler, checkAuthentication)).Methods("POST")
	r.HandleFunc("/api/v1/todos/:id", chainMiddleware(UpdateTodoHandler, checkAuthentication)).Methods("PUT")
	r.HandleFunc("/api/v1/todos/:id", chainMiddleware(DeleteTodoHandler, checkAuthentication)).Methods("DELETE")
	return r
}
