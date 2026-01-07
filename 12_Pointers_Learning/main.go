package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers Learning Class")

	// 1. What is a Pointer?
	// A pointer is a variable that stores the memory address of another variable.

	// var ptr *int
	// Currently The zero value of a pointer is <nil>

	var myNumber int = 23

	// Creating a pointer to myNumber
	// & (Ampersand) operator returns the memory address of a variable
	var ptr *int = &myNumber

	fmt.Println("Value of myNumber:", myNumber)              //Output: Value of myNumber: 23
	fmt.Println("Address of myNumber:", ptr)                 //Output: Address of myNumber: 0xc0000160a8
	fmt.Println("Address of myNumber (using &):", &myNumber) //Output: Address of myNumber (using &): 0xc0000160a8

	// 2. Dereferencing
	// * (Asterisk) operator is used to access the value stored at the pointer's address
	fmt.Println("Value at address stored in ptr:", *ptr) //Output: Value at address stored in ptr: 23

	// 3. Changing value via pointer
	*ptr = *ptr + 2                                 // Increase the value at that address by 2
	fmt.Println("New value of myNumber:", myNumber) //Output: New value of myNumber: 25

	// 4. Pointers and Functions
	// Go is pass-by-value by default. Pointers allow pass-by-reference behavior.
	toChange := 10
	modifyValue(&toChange)
	fmt.Println("Value after modifyValue function:", toChange) //Output: Value after modifyValue function: 20

	Problem1()
	Problem2()
	Problem3()
	Problem4()
}

// Function that takes a pointer to an int
func modifyValue(n *int) {
	*n = *n * 2
}
