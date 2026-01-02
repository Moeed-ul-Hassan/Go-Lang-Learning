package main

//Write a Go program that calculates the duration between two user-provided dates and times.
import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var loc, _ = time.LoadLocation("UTC") // Global location variable for consistency

func inptaking() {
	reader := bufio.NewReader(os.Stdin)
	layout := "2006-01-02 15:04" // Define layout once

	// --- 1. Get Start Date and Time ---
	fmt.Print("Enter Start Date and Time (YYYY-MM-DD HH:MM): ")
	inputStartTimeStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading start date and time:", err)
		return
	}
	inputStartTimeStr = strings.TrimSpace(inputStartTimeStr)

	startTime, err := time.ParseInLocation(layout, inputStartTimeStr, loc)
	if err != nil {
		fmt.Println("Error parsing start date and time:", err)
		return
	}

	// --- 2. Get End Date and Time ---
	fmt.Print("Enter End Date and Time (YYYY-MM-DD HH:MM): ")
	inputEndTimeStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading end date and time:", err)
		return
	}
	inputEndTimeStr = strings.TrimSpace(inputEndTimeStr)

	endTime, err := time.ParseInLocation(layout, inputEndTimeStr, loc)
	if err != nil {
		fmt.Println("Error parsing end date and time:", err)
		return
	}
	// --- 3. Compare Start and End Times ---
	if endTime.Before(startTime) {
		fmt.Println("Error: End time is before start time.")
	} else if endTime.Equal(startTime) {
		fmt.Println("The event starts and ends at the same time.")
	} else {
		fmt.Println("The event is scheduled correctly from", startTime.Format("Monday, January 2, 2006 at 15:04"), "to", endTime.Format("Monday, January 2, 2006 at 15:04"))
	}

	// --- 4. Current Time Comparison ---
	time.Now().In(loc)
	if startTime.Before(time.Now().In(loc)) {
		fmt.Println("The event has already started.")
	} else {
		fmt.Println("The event is yet to start.")
	}
	//---- Final output of start and end time Duaration ----
	duration := endTime.Sub(startTime)
	// Corrected line to include seconds
	fmt.Printf("The event duration is %v hours, %v minutes, and %v seconds.\n",
		int(duration.Hours()), int(duration.Minutes())%60, int(duration.Seconds())%60)

}
func main() {
	inptaking()
}
