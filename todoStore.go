package main

import "github.com/boltdb/bolt"

type TodoStore struct{}

var todoStore TodoStore

// GetTodos returns todos for a user with an email
func (t *TodoStore) GetTodos(email string) ([]byte, error) {
	var todos []byte
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todos"))
		todos = b.Get([]byte("userId"))
		return nil
	})
	if err != nil {
		return []byte(""), err
	}
	return todos, nil
}

// CreateTodo creates a new todo for a user
func (t *TodoStore) CreateTodo(email string, todo []byte) ([]byte, error) {
	resp := []byte("")
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todos"))
		err := b.Put([]byte(email), todo)
		if err != nil {
			return err
		}
		return nil
	})
	return resp, err
}

// InitStores initializes our TodoStore
func InitStores() {
	todoStore = TodoStore{}
}
