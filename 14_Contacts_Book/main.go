// Smart Contact Book
// The Goal: A CLI tool where you can add, search, update, and delete contacts.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func mainmenu() {
	// 1. We need a map to store our contacts (Name -> Phone)
	contacts := make(map[string]string)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n        ================================         ")
		fmt.Println("                CONTACT MANAGER")
		fmt.Println("        ================================         ")
		fmt.Println("        1. Add a contact")
		fmt.Println("        2. Search for a contact")
		fmt.Println("        3. Update a contact")
		fmt.Println("        4. Delete a contact")
		fmt.Println("        5. List all contacts")
		fmt.Println("        6. Exit")
		fmt.Println("        ================================         ")
		fmt.Print("        Select an option: ")

		// 2. Read user input
		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		// 3. Switch on the 'choice' string, not the 'options' map
		switch choice {
		case "1":
			addContact(contacts, reader)
		case "2":
			searchContact(contacts, reader)
		case "3":
			updateContact(contacts, reader)
		case "4":
			deleteContact(contacts, reader)
		case "5":
			listContacts(contacts)
		case "6":
			fmt.Println("Exiting... Goodbye!")
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}

// --- Function Stubs (So the code compiles) ---

func addContact(contacts map[string]string, reader *bufio.Reader) {
	fmt.Println("Add contact logic goes here...")
}

func searchContact(contacts map[string]string, reader *bufio.Reader) {
	fmt.Println("Search contact logic goes here...")
}

func updateContact(contacts map[string]string, reader *bufio.Reader) {
	fmt.Println("Update contact logic goes here...")
}

func deleteContact(contacts map[string]string, reader *bufio.Reader) {
	fmt.Println("Delete contact logic goes here...")
}

func listContacts(contacts map[string]string) {
	fmt.Println("List contacts logic goes here...")
}

func main() {
	mainmenu()
}
