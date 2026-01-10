package main

import "fmt"

func main() {
	// 1. Basic Defer
	// This will print at the very end of main()
	defer fmt.Println("World")
	fmt.Println("Hello")

	fmt.Println("\n--- Multiple Defers (LIFO) ---")
	myDefer()

	fmt.Println("\n--- Real World Example ---")
	// Imagine opening a file here
	fmt.Println("Step 1: Open File")
	defer fmt.Println("Step 3: Close File (Deferred)")
	fmt.Println("Step 2: Write to File")

	fmt.Println("\n--- Practice 1 ---")
	practice()

	fmt.Println("\n--- Practice 2 ---")
	practice2()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		// These are pushed to a stack: 0, 1, 2, 3, 4
		// They will pop out in reverse: 4, 3, 2, 1, 0
		defer fmt.Print(i, " ")
	}
}

// TODO: Practice
// 1. Write a function that uses defer to print "Start" and "End" in the correct order.
// 2. Try to guess the output of a function with 3 defers before running it.

func practice() {
	defer fmt.Println("End")
	fmt.Println("Start")
}

func practice2() {
	defer fmt.Println("End")
	defer fmt.Println("Middle")
	defer fmt.Println("Start")
}
