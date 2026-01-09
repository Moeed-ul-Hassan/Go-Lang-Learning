package main

import "fmt"

type user struct {
	name  string
	age   int
	email string
}

func (u user) greet() {
	fmt.Printf("Hello, my name is %s, I am %d years old, and my email is %s\n", u.name, u.age, u.email)
}

// TODO: Practice Pointer Receivers
// 1. Create a method called 'updateEmail' that takes a newEmail string.
// 2. It MUST use a pointer receiver (*user) so it can change the actual data.
// 3. Inside the method, set u.email = newEmail.
func (u *user) updateEmail(newEmail string) {
	u.email = newEmail
}

// TODO: Another Practice
// 1. Create a method called 'incrementAge' with a pointer receiver.
// 2. It should increase the user's age by 1.
// 3. Call it in main() and print the user's age to verify.

func main() {
	myUser := user{name: "Moeed", age: 22, email: "moeed@google.com"}
	myUser.greet()
	myUser.updateEmail("moeed.new@gmail.com")
	fmt.Println("--- After Update ---")
	myUser.greet()
	myUser.incrementAge()
	fmt.Println("--- After Increment ---")
	myUser.greet()
}

// 4. Test incrementAge here:
func (u *user) incrementAge() {
	u.age++
}
