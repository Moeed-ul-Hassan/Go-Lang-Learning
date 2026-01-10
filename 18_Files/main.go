package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("--- Working with Files in Go ---")

	content := "This is a secret message for Go developers! 🚀"
	fileName := "./mylcfile.txt"

	// 1. Create a file
	file, err := os.Create(fileName)
	checkNilErr(err)

	// 2. Write to the file
	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Printf("Successfully wrote %v characters to the file.\n", length)

	// 3. Close the file (Crucial!)
	// We use defer so it closes at the very end of main
	defer file.Close()

	// 4. Read the file
	readFile(fileName)

	fmt.Println("\n--- Practice Task ---")
	practice()
}

func readFile(filename string) {
	// os.ReadFile returns bytes
	databyte, err := os.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("The content inside the file is:\n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

// TODO: Practice
// 1. Create a new file called "practice.txt".
// 2. Write your name and "I am learning Go" into it.
// 3. Read it back and print it to the console.
func practice() {
	content2 := "Hey I'm Moeed and I am learning Go into it"
	filename2 := "practice.txt"
	file, err := os.Create(filename2)
	checkNilErr(err)
	length, err := io.WriteString(file, content2)
	checkNilErr(err)
	fmt.Printf("Successfully wrote %v characters to the file.\n", length)
	defer file.Close()

	// 3. Read it back
	readFile(filename2)
}
