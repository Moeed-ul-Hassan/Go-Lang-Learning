package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite" // Using SQLite for easy local practice
)

/*
🌟 DAY 1 CHALLENGE: THE CRUD ARCHITECT
Goal: Build a Library Management System (Books Table)

Your Task:
1. Initialize a connection to "library.db".
2. Create a table 'books' with: id, title, author, year.
3. Implement the following functions using PREPARED STATEMENTS.
4. Handle all errors professionally.

Bonus: Add a function to find books by a specific author.
*/

type Book struct {
	ID     int
	Title  string
	Author string
	Year   int
}

func main() {
	// 1. TODO: Open connection to "library.db"
	db, err := sql.Open("sqlite", "library.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 2. TODO: Create 'books' table if it doesn't exist
	createTable(db)

	fmt.Println("--- Day 1: CRUD Mastery Started ---")

	// 3. TODO: Implement and call CRUD functions
	// Example Flow:
	// - Add a book
	// - Read all books
	// - Update a book's year
	// - Delete a book
}

func createTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		author TEXT NOT NULL,
		year INTEGER
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
}

// TODO: Implement AddBook(db *sql.DB, b Book)
// TODO: Implement GetAllBooks(db *sql.DB) ([]Book, error)
// TODO: Implement UpdateBookYear(db *sql.DB, id int, newYear int)
// TODO: Implement DeleteBook(db *sql.DB, id int)
