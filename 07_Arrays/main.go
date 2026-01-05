package main

import "fmt"

func main() {
	fmt.Println("Welcome to Lesson 07: Arrays in Go!")

	// 1. Declaration and Initialization
	// [size]type
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Length of fruit list is: ", len(fruitList))

	// 2. Initializing with values directly
	var vegList = [3]string{"Potato", "Beans", "Mushroom"}
	fmt.Println("Veg list is: ", vegList)
	fmt.Println("Length of veg list is: ", len(vegList))

	// 3. Iterating through an array
	fmt.Println("Iterating through vegList:")
	for i := 0; i < len(vegList); i++ {
		fmt.Printf("Index %d: %s\n", i, vegList[i])
	}

	// 4. Using range for iteration
	fmt.Println("Iterating using range:")
	for index, value := range vegList {
		fmt.Printf("Index %d: %s\n", index, value)
	}
}
