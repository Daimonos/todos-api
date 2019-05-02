package main

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/boltdb/bolt"
	"github.com/google/uuid"
)

// TodoStore is the struct for our todos store
type TodoStore struct{}

var todoStore TodoStore

// Todo is the struct for our todos
type Todo struct {
	ID          string    `json:"id"`
	Todo        string    `json:"todo"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"createdAt"`
	CompletedAt time.Time `json:"completedAt"`
	UserID      string    `json:"user"`
}

// GetTodos returns todos for a user
func (t *TodoStore) GetTodos(email string) ([]Todo, error) {
	todos := []Todo{}
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(email))
		if b == nil {
			return errors.New("bucket does not exist")
		}
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var todo Todo
			err := json.Unmarshal(v, &todo)
			if err != nil {
				return err
			}
			todos = append(todos, todo)
		}
		return nil
	})
	return todos, err
}

// CreateTodo creates a new Todo
func (t *TodoStore) CreateTodo(email string, todo Todo) (Todo, error) {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(email))
		if b == nil {
			b, _ = tx.CreateBucket([]byte(email))
		}
		sid := uuid.New().String()
		todo.ID = sid
		todo.CreatedAt = time.Now()
		todo.Completed = false
		todo.UserID = email
		buf, err := json.Marshal(todo)
		if err != nil {
			return err
		}
		err = b.Put([]byte(sid), buf)
		return err
	})
	return todo, err
}

// UpdateTodo updates a todo for a given user by it's id
func (t *TodoStore) UpdateTodo(email string, id string, todo Todo) (Todo, error) {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(email))
		if b == nil {
			return errors.New("Bucket does not exist for user: " + email)
		}
		if todo.Completed == true {
			todo.CompletedAt = time.Now()
		}
		buf, err := json.Marshal(todo)
		if err != nil {
			return err
		}
		err = b.Put([]byte(id), buf)
		if err != nil {
			return nil
		}
		return nil
	})
	return todo, err
}
