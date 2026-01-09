package main

//I am going to learn one of the Toughest and hardest topics in Go Lang; Interfaces;
// First thing is that  : First Declare the interface then implement it. then Push values and print it...
import (
	"fmt"
	"math"
)

// 1. Define an interface
// An interface is a collection of method signatures. It is a DataType not a function; it is a contract for a type to have certain methods; it is a way to define a type by its behavior.
type shape interface {
	area() float64 //The area contains many methods for circle, rectangle, square etc.
}

// 2. Define structs
type circle struct {
	radius float64 //Here we are defining the circle struct
}

type rectangle struct {
	width, height float64 //Here we are defining the rectangle struct
}

type square struct {
	side float64 //Here we are defining the square struct
}

// 3. Implement the interface methods for each struct
// A type implements an interface by implementing its methods.
// There is no "implements" keyword in Go!
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius //3.14 x radius x radius
}

func (r rectangle) area() float64 {
	return r.width * r.height //width x height
}

func (s square) area() float64 {
	return s.side * s.side //side x side
}

// 4. Use the interface as a type
func printArea(s shape) {
	fmt.Printf("The area is: %.2f\n", s.area())
}

func main() {
	c := circle{radius: 5}
	r := rectangle{width: 10, height: 5}
	sq := square{side: 4}
	sq2 := square2{side: 4}

	fmt.Println("Circle:")
	printArea(c) // circle implements shape

	fmt.Println("Rectangle:")
	printArea(r) // rectangle implements shape

	fmt.Println("Square:")
	printArea(sq) // square implements shape

	fmt.Println("Square2:")
	printArea(sq2) // square2 implements shape
}

// Syntax
// type InterfaceName interface {
//     MethodName(input1 type, input2 type) returnType
//     AnotherMethod() string
// }

// TODO: Practice
// 1. Create a new struct called 'square' with a 'side' field.
// 2. Implement the 'area()' method for 'square'.
// 3. Create a square in main() and call 'printArea()' with it.
type square2 struct {
	side float64
}

func (s square2) area() float64 {
	return s.side * s.side
}
