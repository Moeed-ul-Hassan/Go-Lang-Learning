package main

import (
	"bufio" //buffer input output.
	"fmt"
	"os"
	"strconv" //String conversation.
	"strings"
)

func main() {
	fmt.Println("Welcome to our Pizza App..!")
	fmt.Println("Please enter the pizza size you want to order (e.g., 12 for 12 inches):")
	reader := bufio.NewReader(os.Stdin)
	sizeInput, _ := reader.ReadString('\n')
	trimmedSizeInput := strings.TrimSpace(sizeInput)
	pizzaSize, err := strconv.ParseFloat(trimmedSizeInput, 64)
	if err != nil {
		fmt.Println("Invalid input. Please enter a numeric value for the pizza size.")
		return
	}

	fmt.Printf("Thanks for Ordering! You have entered %.2f inches pizza size.\n", pizzaSize)
}
