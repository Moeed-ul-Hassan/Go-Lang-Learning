package main

import "fmt"

// 2. Define a function 'increment(val *int)'
func increment(val *int) {
	// 3. Inside the 'increment' function, dereference the pointer and increment the value it points to by 1.
	*val = *val + 1 // Or simply *val++
}

func main() {
	// 1. Declares an integer variable 'count' and initializes it to 5.
	count := 5
	fmt.Println("Initial count:", count) // 5. Print before

	// 4. In main, call the 'increment' function, passing the address of 'count'.
	increment(&count)

	fmt.Println("Count after increment:", count) // 5. Print after
}
