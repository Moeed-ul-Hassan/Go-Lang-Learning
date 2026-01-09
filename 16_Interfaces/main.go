package main

import (
	"fmt"
	"math"
)

// 1. Define an interface
// An interface is a collection of method signatures.
type shape interface {
	area() float64
}

// 2. Define structs
type circle struct {
	radius float64
}

type rectangle struct {
	width, height float64
}

// 3. Implement the interface methods for each struct
// A type implements an interface by implementing its methods.
// There is no "implements" keyword in Go!
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

// 4. Use the interface as a type
func printArea(s shape) {
	fmt.Printf("The area is: %.2f\n", s.area())
}

func main() {
	c := circle{radius: 5}
	r := rectangle{width: 10, height: 5}

	fmt.Println("Circle:")
	printArea(c) // circle implements shape

	fmt.Println("Rectangle:")
	printArea(r) // rectangle implements shape
}

// TODO: Practice
// 1. Create a new struct called 'square' with a 'side' field.
// 2. Implement the 'area()' method for 'square'.
// 3. Create a square in main() and call 'printArea()' with it.
