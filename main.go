package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Hello, Todos API")
	InitDB("todos-dev")
	r := NewRouter()
	log.Panic(http.ListenAndServe(":8081", r))
}
