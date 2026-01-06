package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")

	// Maps are key-value pairs
	// make(map[KeyType]ValueType)

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages:", languages)
	fmt.Println("JS shorts for:", languages["JS"])

	// Delete a key
	delete(languages, "RB")
	fmt.Println("List of all languages after delete:", languages)

	// Iterating over map
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
