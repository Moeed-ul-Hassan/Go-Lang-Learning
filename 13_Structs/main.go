package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Email  string
	Active bool
}

func main() {
	user := User{
		Name:   "John Doe",
		Age:    30,
		Email:  "john.doe@example.com",
		Active: true,
	}
	fmt.Println(user)
}
