package main

import (
	"fmt"
	"time"
)

func sepcifictime() {
	fmt.Println("Current Time:", time.Now())
	presentTime := time.Now()
	fmt.Println(presentTime.Format("2006-01-02 15:04:05"))
	//extracing specific time
	createdDate := time.Date(2026, time.September, 20, 12, 0, 0, 0, time.UTC)
	fmt.Println("Created Date:", createdDate, time.Now().Local().Format("2006-01-02 15:04:05"))
}
func errandtime() {
	timeString := "2023-10-15 14:30:00"
	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeString)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	} else {
		fmt.Println("Parsed Time:", parsedTime)
	}

}
func main() {
	sepcifictime()
	errandtime()
}
