package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Expense struct to keep name and amount together
type Expense struct {
	Name   string
	Amount float64
}

// Global slice to store all expenses
var expenses []Expense

func mainmenu() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Personal Expense Tracker ---")
		fmt.Println("1. Add Expense")
		fmt.Println("2. View Expenses")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Invalid Choice")
			continue
		}

		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			addExpense(reader)
		case "2":
			viewExpenses()
		case "3":
			Exit()
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func addExpense(reader *bufio.Reader) {
	fmt.Print("Enter expense name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter expense amount: ")
	amountStr, _ := reader.ReadString('\n')
	amountStr = strings.TrimSpace(amountStr)

	// Convert the string input to a number (float64)
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		fmt.Println("Invalid amount! Please enter a number.")
		return
	}

	// Professional Validation: Amount must be positive
	if amount <= 0 {
		fmt.Println("Error: Expense amount must be a positive number.")
		return
	}

	// Create a new Expense struct instance
	newExpense := Expense{
		Name:   name,
		Amount: amount,
	}

	// Add the struct to our slice
	expenses = append(expenses, newExpense)

	fmt.Println("Expense added successfully!")
}

func viewExpenses() {
	if len(expenses) == 0 {
		fmt.Println("No expenses found.")
		return
	}

	fmt.Println("\n--- Your Expenses ---")
	var total float64
	for i, exp := range expenses {
		fmt.Printf("%d. %s: $%.2f\n", i+1, exp.Name, exp.Amount)
		total += exp.Amount
	}
	fmt.Printf("----------------------\nTotal: $%.2f\n", total)
}

func Exit() {
	fmt.Println("Good Bye! Expense Tracker")
	os.Exit(0)
}

func main() {
	mainmenu()
}
