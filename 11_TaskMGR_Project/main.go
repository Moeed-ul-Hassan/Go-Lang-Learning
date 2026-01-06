package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func TM() {
	fmt.Println("Task Manager")

	tasklist := make([]string, 0)
	reader := bufio.NewReader(os.Stdin)

	for {
		//---------------Game Menu---------------------------------------------
		fmt.Println("\n--------------------------------")
		fmt.Println("Press 'a' to add a task")
		fmt.Println("Press 'l' to list all tasks")
		fmt.Println("Press 'r' to remove a task")
		fmt.Println("Press 'q' to quit")
		fmt.Print("Enter option: ")
		//----------------------------------------------------------------------
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input")
			continue
		}
		input = strings.TrimSpace(input)
		//----------------------------------------------------------------------
		switch input {
		case "a":
			fmt.Println("Enter Task to Add:")
			taskName, _ := reader.ReadString('\n')
			taskName = strings.TrimSpace(taskName)
			if taskName != "" {
				tasklist = append(tasklist, taskName)
				fmt.Println("Task added:", taskName)
			}
			//----------------------------------------------------------------------
		case "l":
			fmt.Println("List all tasks:")
			if len(tasklist) == 0 {
				fmt.Println("No tasks found.")
			} else {
				for i, task := range tasklist {
					fmt.Printf("%d. %s\n", i+1, task)
				}
			}
			//----------------------------------------------------------------------
		case "r":
			fmt.Println("Enter Task to Remove:")
			taskToRemove, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input")
				continue
			}
			taskToRemove = strings.TrimSpace(taskToRemove)

			index := -1
			for i, task := range tasklist {
				if task == taskToRemove {
					index = i
					break
				}
			}
			//------------------------------------------------------------------------
			if index != -1 {
				tasklist = append(tasklist[:index], tasklist[index+1:]...)
				fmt.Println("Task removed:", taskToRemove)
			} else {
				fmt.Println("Task not found.")
			}
		case "q":
			fmt.Println("Quitting...")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}
func main() {
	TM()
}
