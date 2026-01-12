package main

import (
	"encoding/json" //Compulsory to Import while working with JSON.
	"fmt"
	"log"      //Compulsory to Import while logging.
	"net/http" //Compulsory to Import while making a web server.
	"strings"
)

// Message represents a simple JSON payload for POST /echo
type Message struct {
	Text string `json:"text"`
}

func mainserver() {
	//It is Compulsory to Write This as they are routes for our server
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/hello/", helloHandler)  //it will handle /hello/{name}
	mux.HandleFunc("/search", searchHandler) //It will handle /search?q=term
	mux.HandleFunc("/echo", echoHandler)     //It will handle POST /echo expects JSON {"text":"..."} and echoes it back
	mux.HandleFunc("/submit", submitHandler) //It will handle POST /submit expects form data and returns it as JSON
	mux.HandleFunc("/data", dataHandler)     //It will handle GET /data returns a static JSON array of courses (example data)
	mux.HandleFunc("/about", aboutHandler)

	fmt.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Week 3 – Hello World")
}

// /hello/{name}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Trim the prefix "/hello/" to get the name
	name := strings.TrimPrefix(r.URL.Path, "/hello/")
	if name == "" {
		http.Error(w, "name not provided", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

// /search?q=term returns JSON {"query":"term"}
func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	resp := map[string]string{"query": query}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// POST /echo expects JSON {"text":"..."} and echoes it back
func echoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST method required", http.StatusMethodNotAllowed)
		return
	}
	var msg Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

// POST /submit expects form data and returns it as JSON
func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST method required", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "cannot parse form", http.StatusBadRequest)
		return
	}
	resp := make(map[string]string)
	for key, values := range r.PostForm {
		if len(values) > 0 {
			resp[key] = values[0]
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// /data returns a static JSON array of courses (example data)
func dataHandler(w http.ResponseWriter, r *http.Request) {
	type Course struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	}
	courses := []Course{{ID: 1, Title: "Intro to Go"}, {ID: 2, Title: "Web & JSON Basics"}}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
func main() {
	problem1()
	mainserver()
}
