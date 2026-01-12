package main

import (
	"fmt"
	"log"
	"net/http"
)

// Problem 1 : The "About" Page

func problem1() {
	mux := http.NewServeMux()
	mux.HandleFunc("/about", aboutHandler)

	fmt.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Write your code here
	w.Write([]byte("This is my first Go Web Server!"))
}
