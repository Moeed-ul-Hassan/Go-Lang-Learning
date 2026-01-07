package main

import "fmt"

func Problem1() {
	// ============================================================
	// Problem 1: The Basics (Easy)
	// ============================================================
	// 1. Declare an integer variable `score` with value 100.
	// 2. Create a pointer `pScore` that points to `score`.
	// 3. Print the value of `score` using the pointer `pScore`.
	// 4. Change the value of `score` to 200 using ONLY the pointer `pScore`.
	// 5. Print `score` again to verify the change.

	fmt.Println("--- Problem 1 ---")
	var score int = 100
	var pScore *int = &score
	fmt.Println(*pScore) // Output 100
	*pScore = 200
	fmt.Println(score) // Output 200
}

// ============================================================
// Problem 2: The Swap (Medium)
// ============================================================
// 1. Write a function `swap(a, b *int)` that swaps the values of two integers.
// 2. In main, declare x = 15, y = 30.
// 3. Call `swap` passing the addresses of x and y.
// 4. Print x and y to verify they have been swapped.

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func Problem2() {
	fmt.Println("\n--- Problem 2 ---")
	x, y := 15, 30
	fmt.Printf("Before Swap: x=%d, y=%d\n", x, y)
	swap(&x, &y)
	fmt.Printf("After Swap: x=%d, y=%d\n", x, y)
}

// ============================================================
// Problem 3: Pointer to Pointer (Medium)
// ============================================================
// 1. Declare an integer `val` = 50.
// 2. Create a pointer `p1` to `val`.
// 3. Create a pointer `p2` to `p1` (a pointer to a pointer).
// 4. Print `val` using `p2`.
// 5. Change `val` to 100 using ONLY `p2`.
// 6. Print `val` to verify.

func Problem3() {
	fmt.Println("\n--- Problem 3 ---")
	val := 50
	p1 := &val
	p2 := &p1
	// To print the value, we need to dereference twice: **p2
	fmt.Println("Value using p2:", **p2)

	// To change the value, we also dereference twice
	**p2 = 100
	fmt.Println("New value of val:", val)
}

// ============================================================
// Problem 4: Nil Pointer Check (Easy)
// ============================================================
// 1. Write a function `printIfSafe(ptr *int)` that prints the value if not nil,
//    otherwise prints "Pointer is nil".
// 2. Call it with a valid pointer and then with a nil pointer.

func printIfSafe(ptr *int) {
	if ptr != nil {
		fmt.Println("Value:", *ptr)
	} else {
		fmt.Println("Pointer is nil")
	}
}

func Problem4() {
	fmt.Println("\n--- Problem 4 ---")

	// 1. Valid pointer
	num := 42
	printIfSafe(&num)

	// 2. Nil pointer
	var nilPtr *int
	printIfSafe(nilPtr)
}
