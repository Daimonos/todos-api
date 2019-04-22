package main

// User is the user struct we expect to receive from our User Service
type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}
