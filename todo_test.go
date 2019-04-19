package main

import (
	"log"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	InitDB("test.db")
	todoStore = TodoStore{}
	retCode := m.Run()
	db.Close()
	os.Remove("test.db")
	os.Exit(retCode)
}

func TestTodoErrorNoBucket(t *testing.T) {
	_, err := todoStore.GetTodos("test@domain.com")
	if err == nil {
		t.Fatalf(err.Error())
	}
}

func TestTodoCreateSuccess(t *testing.T) {
	todo := Todo{
		Todo:      "Test",
		CreatedAt: time.Now(),
		Completed: false,
	}
	_, err := todoStore.CreateTodo("test@domain.com", todo)
	if err != nil {
		t.Fatalf(err.Error())
	}
	todos, err := todoStore.GetTodos("test@domain.com")
	if err != nil {
		log.Fatalf(err.Error())
	}
	if len(todos) != 1 {
		log.Fatalf("Expected exactly 1 Todo")
	}
}

func TestTodoCreateMultipleSuccess(t *testing.T) {
	todo := Todo{
		Todo:      "Test",
		CreatedAt: time.Now(),
		Completed: false,
	}
	_, err := todoStore.CreateTodo("test@domain.com", todo)
	if err != nil {
		t.Fatalf(err.Error())
	}
	todos, err := todoStore.GetTodos("test@domain.com")
	if err != nil {
		log.Fatalf(err.Error())
	}
	if len(todos) != 2 {
		log.Fatalf("Expected exactly 2 Todos")
	}
}

func TestTodoGetWithMultipleUsersOnlyGetOne(t *testing.T) {
	todo := Todo{
		Todo:      "Test",
		CreatedAt: time.Now(),
		Completed: false,
	}
	_, err := todoStore.CreateTodo("testa@domain.com", todo)
	_, err = todoStore.CreateTodo("testb@domain.com", todo)
	if err != nil {
		t.Fatalf(err.Error())
	}
	var todos []Todo
	todos, err = todoStore.GetTodos("testa@domain.com")
	if len(todos) > 1 {
		t.Fatalf("Expected only 1 todo to be returned in database with 2 users")
	}
}
