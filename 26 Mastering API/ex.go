package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func problem1() {
	fmt.Println("json.marshal is is used to write data in JSON from Main file")
	fmt.Println("I Don't know")
	fmt.Println("By default, Content type is set to text/plain")
}
func problem2() {
	fmt.Println("--Problem 2--")
	type Product struct {
		ProductID int     `json:"product_id"`
		Name      string  `json:"name"`
		Price     float64 `json:"price"`
	}
	product := Product{
		ProductID: 1,
		Name:      "Laptop",
		Price:     1000.0,
	}
	json.NewEncoder(os.Stdout).Encode(product)
}
