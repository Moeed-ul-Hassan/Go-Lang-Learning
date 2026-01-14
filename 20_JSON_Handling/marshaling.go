package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON Marshaling in Go!") //Marshaling is the process of taking Go data and turning it into JSON...
	encodeJson()
}

func encodeJson() {
	courses := []Course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "password123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "angular123", nil},
	}

	// Package this data as JSON data
	finalJson, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

	// Writing to external file
	err = os.WriteFile("main.json", finalJson, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Data successfully saved to main.json")
}
