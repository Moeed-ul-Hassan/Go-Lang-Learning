package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to Go Modules Learning! 🏎️")

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	fmt.Println("Server starting on port 4000...")
	log.Fatal(http.ListenAndServe(":4000", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Go Modules!</h1>"))
}
