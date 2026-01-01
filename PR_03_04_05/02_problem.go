// Problem 2
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var loc, _ = time.LoadLocation("UTC") // Global location variable

func entsched() {
	fmt.Print("Enter an event date (YYYY-MM-DD): ")
	dateInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading date input, try again:", err)
		return
	}
	dateInput = strings.TrimSpace(dateInput)

	fmt.Print("Enter an event time (HH:MM): ")
	timeInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading time input, try again:", err)
		return
	}
	timeInput = strings.TrimSpace(timeInput)

	// Combine date and time strings
	dateTimeString := fmt.Sprintf("%s %s", dateInput, timeInput)

	// Define the layout for parsing (must match the combined string format)
	layout := "2006-01-02 15:04"

	// Parse the combined string into a time.Time object
	eventTime, parseErr := time.ParseInLocation(layout, dateTimeString, loc)
	if parseErr != nil {
		fmt.Println("Invalid date-time format. Please use YYYY-MM-DD HH:MM.", parseErr)
		return
	}

	now := time.Now().In(loc) // Get current time in the specified location

	// Compare the eventTime with the current time
	if eventTime.Before(now) {
		fmt.Println("This event is in the past.")
	} else if eventTime.After(now) {
		fmt.Println("This event is in the future.")
	} else {
		fmt.Println("This event is happening right now.")
	}

	// Corrected Println to fmt.Println and used eventTime
	fmt.Println("Event scheduled for", eventTime.Format("Monday, January 2, 2006 at 15:04"))
}

// You might want to add a main function to call entsched for testing
/*
func main() {
    entsched()
}
*/
