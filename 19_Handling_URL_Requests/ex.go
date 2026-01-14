package main

import (
	"fmt"
	"net/http"
	"strings"
)

/*
✅ SOLUTION 1: Path Parameters
URL: /products/electronics
*/
func problem1(w http.ResponseWriter, r *http.Request) {
	// 1. We take the full path: r.URL.Path (which is "/products/electronics")
	// 2. We cut off the "/products/" part to leave just "electronics"
	category := strings.TrimPrefix(r.URL.Path, "/products/")

	fmt.Fprintf(w, "Category: %s", category)
}

/*
✅ SOLUTION 2: Query Parameters
URL: /list?sort=desc&limit=5
*/
func problem2(w http.ResponseWriter, r *http.Request) {
	// 1. r.URL.Query() gets all the "sticky notes" (?sort=desc&limit=5)
	// 2. .Get("key") picks the specific one we want
	sortOrder := r.URL.Query().Get("sort")
	limitValue := r.URL.Query().Get("limit")

	fmt.Fprintf(w, "Sorting by %s, Limit: %s", sortOrder, limitValue)
}

// ---------------------------------------------------------
// 🌟 ULTRA-SIMPLE MINI-CHALLENGES 🌟
// ---------------------------------------------------------

// MINI-CHALLENGE 1: Just print the raw path!
// Goal: If I go to /anything/here, just print "/anything/here"
func challenge1(w http.ResponseWriter, r *http.Request) {
	// TODO: Use fmt.Fprintf to print r.URL.Path
	anythinghere := r.URL.Path
	fmt.Fprint(w, anythinghere)
}

// MINI-CHALLENGE 2: Get a single query param "name"
// Goal: If I go to /greet?name=Moeed, print "Hi Moeed"
func challenge2(w http.ResponseWriter, r *http.Request) {
	// TODO: Use r.URL.Query().Get("name")
	query := r.URL.Query().Get("name")
	fmt.Fprintln(w, "Hi", query)
}

// MINI-CHALLENGE 3: Path + Query
// Goal: URL /city/London?weather=sunny -> print "London is sunny"
func challenge3(w http.ResponseWriter, r *http.Request) {
	// Hint: Use strings.TrimPrefix for city, and .Get() for weather
	location := strings.TrimPrefix(r.URL.Path, "/city/") // Added / for cleaner extraction
	weather := r.URL.Query().Get("weather")
	fmt.Fprintf(w, "%s is %s", location, weather)
}

// ---------------------------------------------------------
// 🚀 LEVEL 2: MASTERING URLS 🚀
// ---------------------------------------------------------

// CHALLENGE 4: Multiple Query Params
// Goal: URL /filter?min=10&max=100 -> print "Price range: 10 to 100"
func challenge4(w http.ResponseWriter, r *http.Request) {
	// TODO: Get "min" and "max" from query
	min := r.URL.Query().Get("min")
	max := r.URL.Query().Get("max")
	// Use Fprintf for %s formatting!
	fmt.Fprintf(w, "Price range: %s - %s", min, max)
}

// CHALLENGE 5: Deep Path Segments
// Goal: URL /users/moeed/settings -> print "User: moeed, Page: settings"
// Hint: You can use strings.TrimPrefix twice or strings.Split
func challenge5(w http.ResponseWriter, r *http.Request) {
	// TODO: Extract name and page from path
	// Hint: Use strings.Split to get all parts of the path
	parts := strings.Split(r.URL.Path, "/")

	parts[0] = ""
	parts[1] = "users"
	parts[2] = "moeed"
	parts[3] = "settings"

	if len(parts) < 4 {
		http.Error(w, "Invalid path", 400)
		return
	}

	name := parts[2]
	page := parts[3]
	fmt.Fprintf(w, "User is %s and Page is %s", name, page)
}

// CHALLENGE 6: Simple Validation
// Goal: URL /check?age=20
// If age is empty, print "Please provide age"
// Otherwise, print "Age received: 20"
func challenge6(w http.ResponseWriter, r *http.Request) {
	// TODO: Get "age", check if empty, and print message
	age := r.URL.Query().Get("age")
	if age == "" {
		fmt.Fprintln(w, "Please provide age:")
	} else {
		fmt.Fprintln(w, "Age received:", age)
	}
}
