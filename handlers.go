package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/context"
)

// GetTodosHandler is a handler for managing Todos for a user
func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if u := context.Get(r, "UserId"); u != nil {
		WriteError(w, http.StatusUnauthorized, errors.New("Not Authorized"))
		return
	}
	todos, err := todoStore.GetTodos(user.Email)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return
	}
	WriteJSON(w, http.StatusOK, todos)
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
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(err.Error()))
	return
}
