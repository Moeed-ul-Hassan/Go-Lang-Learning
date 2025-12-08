// All Variables Type are Described Below
package main

import "fmt"

func main() {
	// Integer Types
	var a int = 10      // only nmber without decimal, upto 10 digits
	var b int8 = 20     // -128 to 127
	var c int16 = 300   // -32,768 to 32,767
	var d int32 = 4000  // -2,147,483,648 to 2,147,483,647
	var e int64 = 50000 // -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
	fmt.Println(a)      //printing the values of a
	fmt.Println(b)      //printing the values of b
	fmt.Println(c)      //printing the values of c
	fmt.Println(d)      //printing the values of d
	fmt.Println(e)      //printing the values of e
}

func main2() {
	//Float Data Types are described below
	var f float32 = 5.678    // upto 7 decimal digits
	var g float64 = 9.101112 // upto 15 decimal digits
	// Here is what if you print more than the specified decimal digits
	var f1 float32 = 5.6789123456       // will print 5.678912
	var g1 float64 = 9.1011121314151617 // will print 9.101112131415162

	// Print values
	fmt.Println(f)  //printing the values of f
	fmt.Println(g)  //printing the values of g
	fmt.Println(f1) //printing the values of f1
	fmt.Println(g1) //printing the values of g1

	x
}

func main3() {
	//Other Data Types are described below
	var h bool = true           // boolean type
	var i string = "Hello, Go!" // string type
	var j byte = 'A'            // byte type
	var k rune = '₹'            // rune type
	fmt.Println(h)              //pr
	// inting the values of h
	fmt.Println(i) //printing the values of i
	fmt.Println(j) //printing the values of j
	fmt.Println(k) //printing the values of k
}
func main4() {
	// Complex Data Types are described below
	var l complex64 = 1 + 2i  // complex number with float32 real and imaginary parts
	var m complex128 = 3 + 4i // complex number with float64 real and imaginary parts
	fmt.Println(l)            //printing the values of l
	fmt.Println(m)            //printing the values of m
}
func main5() {
	// Unsigned Integer Types are described below
	var n uint = 10    // only positive numbers without decimal
	var o uint8 = 20   // 0 to 255
	var p uint16 = 300 // 0 to 65,535
	var q uint32 = 4000
	var r uint64 = 50000
	fmt.Println(n) //printing the values of n
	fmt.Println(o) //printing the values of o
	fmt.Println(p) //printing the values of p
	fmt.Println(q) //printing the values of q
	fmt.Println(r) //printing the values of r
}
func main6() {
	// Alias Data Types are described below
	var s byte = 65  // alias for uint8
	var t rune = '₹' // alias for int32
	fmt.Println(s)   //printing the values of s
	fmt.Println(t)   //printing the values of t
}

// Note: Only one main function can be executed at a time. To test other main functions, rename them to main and run the program.
//written by Moeed ul Hassan... Date : 08-12-2025;
// in  next files we will learn about constants and operators in golang.
//Allah Hafiz and All the best... :)
