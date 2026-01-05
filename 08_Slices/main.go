package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Lesson 08: Slices in Go!")

	// 1. Declaration and Initialization
	// Slices are like dynamic arrays. You don't specify the size.
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	// 2. Appending to a slice
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println("Fruit list after append: ", fruitList)

	// 3. Slicing operations [start:end]
	// start is inclusive, end is exclusive
	fmt.Println("Slicing [1:3]: ", fruitList[1:3])
	fmt.Println("Slicing [1:]: ", fruitList[1:])
	fmt.Println("Slicing [:3]: ", fruitList[:3])

	// 4. Using make() to create a slice
	// make([]type, length, capacity)
	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 777 // This would cause an error because length is 4

	// However, append() works because it reallocates memory if needed
	highScores = append(highScores, 555, 666, 321)

	fmt.Println("High scores: ", highScores)

	// 5. Sorting a slice
	fmt.Println("Is sorted: ", sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println("Sorted high scores: ", highScores)
	fmt.Println("Is sorted now: ", sort.IntsAreSorted(highScores))

	// 6. Removing a value from a slice based on index
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println("Courses: ", courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Courses after removing 'swift': ", courses)
}
