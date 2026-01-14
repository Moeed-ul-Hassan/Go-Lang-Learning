package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func greeting() {
	fmt.Println("Hello! welcome to our Code Academy!!\n", "Here are Course which are availible!!")
}

type courselist struct {
	Name        string `json:courseName`
	Price       int
	Discription string `json:courseDiscription`
	Rating      int
}

func encodeJsoning() {
	courses := []courselist{
		{"Go Lang", 50, "A breif course and practical Course on Go Lang Cloud.", 4},
		{"Python", 45, "A breif course and practical Course on Python Machine Learning.", 4},
		{"Web Dev", 10, "A breif and practical Course on Full Stack Web Development.", 3},
	}
	finalJsoning, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJsoning)

	err = os.WriteFile("main2.json", finalJsoning, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully written All Data in JSON")
}
func main() {
	greeting()
	encodeJsoning()
}
