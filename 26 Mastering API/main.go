package main

import (
	"encoding/json" //JSON Handling
	"fmt"
	"log"
	"net/http" //Handling HTTP requests
)

// Article defines the structure of a single article
type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

// Articles is a slice of Article
type Articles []Article

var allArticles Articles

func serverHomepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the homepage!")
}

func articlesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: articlesHandler")

	// Set the header so the client knows it's receiving JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode our allArticles slice into JSON and send it
	json.NewEncoder(w).Encode(allArticles)
}

func main() {
	// Initialize data
	allArticles = Articles{
		{
			Title:   "Article 1",
			Content: "This is the content of article 1",
			Author:  "Author 1",
		},
		{
			Title:   "Article 2",
			Content: "This is the content of article 2",
			Author:  "Author 2",
		},
	}

	// Register routes
	http.HandleFunc("/", serverHomepage)
	http.HandleFunc("/articles", articlesHandler)

	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
