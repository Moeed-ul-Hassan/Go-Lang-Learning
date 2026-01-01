package main

// := is an walrus operator
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin) //stdin means standard input. It is a way to read input from the user via the terminal or command line.
	fmt.Println("Enter text:")
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered:", input)
}
