package main

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/todos", chainMiddleware(GetTodosHandler, withLogging, withTracing, checkAuthentication)).Methods("GET")
	r.HandleFunc("/api/v1/todos", chainMiddleware(CreateTodoHandler, withLogging, withTracing, checkAuthentication)).Methods("POST")
	return r
}
