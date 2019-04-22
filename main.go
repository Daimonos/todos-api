package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Todos API Starting Up")
	port := getEnv("TODOS_PORT", ":8080")
	dbname := getEnv("TODOS_DBNAME", "todos-dev")
	InitDB(dbname)
	r := NewRouter()
	log.Println("App listneing onport: " + port)
	log.Println("Db Name: " + dbname)
	log.Panic(http.ListenAndServe(port, r))
}

func getEnv(variable, fallback string) string {
	log.Println("Looing for environment variable: " + variable)
	if value, ok := os.LookupEnv(variable); ok {
		return value
	}
	log.Println("Returning default value: " + fallback)
	return fallback
}
