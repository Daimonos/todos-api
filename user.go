package main

type User struct {
	Email    string `json:"email"`
	FullName string `password:"password"`
}
