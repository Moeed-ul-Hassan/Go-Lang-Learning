package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Todo List App!")

	// TODO: Create a slice to store tasks

	tasks := make([]string, 4)
	tasks[0] = "Programming"
	tasks[1] = "Making Coffee"
	tasks[2] = "Going to College"
	tasks[3] = "Back to Home"

	fmt.Println("Your Current Tasks are: ", tasks)

	var userChoice string
	fmt.Println("Press 'a' to add a task")
	fmt.Scanln(&userChoice)

	if userChoice == "a" {
		var newTask string
		fmt.Println("Enter the task to Add:")
		fmt.Scanln(&newTask)
		tasks = append(tasks, newTask)
		fmt.Println("Task added!")
	} else {
		fmt.Println("Invalid Input!")
	}
	fmt.Println(tasks[3])
	// TODO: Create a loop that keeps asking the user for input
	// for i:= 0; i < 1999; i++{
	//  fmt.Println("Enter a Charcter:")
	//  fmt.Scanln(&tasks[i])
	// }
	// TODO: Implement options: 1. Add, 2. List, 3. Delete, 4. Exit
	var options string

	for {
		fmt.Println("\nEnter your option (add, list, delete, exit):")
		fmt.Scanln(&options)

		switch options {
		case "add":
			var newTask string
			fmt.Println("Enter the task:")
			fmt.Scanln(&newTask)
			tasks = append(tasks, newTask)
			fmt.Println("Task added!")

		case "list":
			fmt.Println("Current Tasks:")
			for i, task := range tasks {
				fmt.Printf("%d. %s\n", i, task)
			}

		case "delete":
			var index int
			fmt.Println("Enter the task number to delete:")
			fmt.Scanln(&index)

			if index >= 0 && index < len(tasks) {
				tasks = append(tasks[:index], tasks[index+1:]...)
				fmt.Println("Task deleted!")
			} else {
				fmt.Println("Invalid task number")
			}

		case "exit":
			fmt.Println("Thank you for using the Todo List App!")
			return

		default:
			fmt.Println("Invalid option")
		}
	}
}
