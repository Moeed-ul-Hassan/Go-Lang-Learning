package main

//Answer is "If you put everything in main.go, it would quickly become hundreds of lines long."
import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/", redirectHandler)

	fmt.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error while starting server: %v\n", err)
	}
}
func shortenHandler(w http.ResponseWriter, r *http.Request) {

}
func redirectHandler(w http.ResponseWriter, r *http.Request) {

}
