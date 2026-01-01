package main

import (
	"bufio"
	"fmt"
	"os"
	"strings" // Import the strings package for TrimSpace
	"time"
)

func prtimgen() {
	fmt.Println("Personalized Greeting with Current Time Generator")
	var (
		username string
		dantim   string // Variable declared as dantim
	)
	fmt.Print("Enter your name: ")
	// Corrected bufio usage and ReadString method call
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		// If there's an error reading the name, still set dantim
		dantim = time.Now().Format("2006-01-02 15:04:05")
		// username will remain an empty string in this case
	} else {
		// Assign the trimmed input to username
		username = strings.TrimSpace(input)
		// Assign the current time to dantim
		dantim = time.Now().Format("2006-01-02 15:04:05")
	}
	fmt.Printf("Hello, %s! Today is %s.\n", username, dantim)
}
func main() {
	prtimgen()
}
