package main

import "fmt"

type User struct {
	Name  string
	Email string
	Age   int
}

// Value Receiver: Works on a copy
func (u User) Greet() {
	fmt.Printf("Hello, my name is %s and my email is %s\n", u.Name, u.Email)
}

// TODO: Add a Pointer Receiver method called UpdateEmail here

func main() {
	myUser := User{"Moeed", "moeed@example.com", 22}
	myUser.Greet()

	// After you implement UpdateEmail, call it here:
	// myUser.UpdateEmail("newemail@example.com")
	// myUser.Greet()
}
