package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/context"
)

type Payload struct {
	ErrorCode    int
	ErrorMessage string
}

// GetTodosHandler is a handler for managing Todos for a user
func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	u := context.Get(r, "User")
	if u == nil {
		WriteError(w, http.StatusUnauthorized, errors.New("Not Authorized"))
		return
	}
	user := u.(User)
	todos, err := todoStore.GetTodos(user.Email)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return
	}
	WriteJSON(w, http.StatusOK, todos)
}

// CreateTodoHandler handles the creation of todos for a user
func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	u := context.Get(r, "User")
	if u == nil {
		WriteError(w, http.StatusUnauthorized, errors.New("Not Authorized"))
		return
	}
	log.Println(u)
	user := u.(User)
	log.Println("Creating TODO for user: " + user.Email)
	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		log.Println("Error decoding TODO")
		WriteError(w, http.StatusInternalServerError, err)
		return
	}
	todo, err := todoStore.CreateTodo(user.Email, todo)
	if err != nil {
		log.Println("GOT ERROR")
		WriteError(w, http.StatusInternalServerError, err)
		return
	}
	WriteJSON(w, http.StatusOK, todo)
}

// WriteJSON is a helper function for writing JSON content
func WriteJSON(w http.ResponseWriter, code int, payload interface{}) {
	resp, err := json.Marshal(payload)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, errors.New("Error encoding payload"))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)
	return
}

// WriteError is a helper function for writing Error content
func WriteError(w http.ResponseWriter, code int, err error) {
	w.Header().Add("Content-Type", "text")
	w.WriteHeader(code)
	w.Write([]byte(err.Error()))
	return
}
