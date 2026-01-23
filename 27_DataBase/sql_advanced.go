package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

/*
"This is an AI Code"
Database Mastery: Advanced SQL (sqlx & sqlc)
This file demonstrates the use of sqlx for easier struct mapping and named queries.
It also explains the concept of sqlc.
*/

// Student struct with db tags for sqlx mapping
type Student struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func main() {
	// 1. Use sqlx.Open() instead of sql.Open()
	// sqlx.DB is a wrapper around sql.DB that provides extra functionality.
	db, err := sqlx.Open("sqlite", "learning.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Connected to SQLite using sqlx!")

	// 2. Easier handling of structs: Use db.Select() or db.Get()
	// db.Select() retrieves multiple rows and maps them to a slice of structs.
	var students []Student
	err = db.Select(&students, "SELECT * FROM students")
	if err != nil {
		log.Fatal("Error selecting students:", err)
	}

	fmt.Println("Students (retrieved via sqlx.Select):")
	for _, s := range students {
		fmt.Printf("- ID: %d, Name: %s, Age: %d\n", s.ID, s.Name, s.Age)
	}

	// db.Get() retrieves a single row and maps it to a struct.
	var singleStudent Student
	err = db.Get(&singleStudent, "SELECT * FROM students WHERE name = ?", "Moeed")
	if err != nil {
		log.Fatal("Error getting student:", err)
	}
	fmt.Printf("Single Student (retrieved via sqlx.Get): %s, Age: %d\n", singleStudent.Name, singleStudent.Age)

	// 3. Named queries: Use db.NamedExec()
	// This allows you to use struct field names or map keys in your queries.
	newStudent := Student{Name: "Charlie", Age: 21}
	_, err = db.NamedExec("INSERT INTO students (name, age) VALUES (:name, :age)", newStudent)
	if err != nil {
		log.Fatal("Error inserting via named query:", err)
	}
	fmt.Println("Inserted Charlie using a named query.")

	// --- Conceptual Overview: sqlc ---
	fmt.Println("\n--- Conceptual Overview: sqlc ---")
	fmt.Println("sqlc is a tool that generates type-safe Go code from your SQL queries.")
	fmt.Println("How it works:")
	fmt.Println("1. You write your database schema (DDL) in a .sql file.")
	fmt.Println("2. You write your SQL queries in another .sql file.")
	fmt.Println("3. You run 'sqlc generate'.")
	fmt.Println("4. sqlc parses the SQL and generates Go interfaces/structs that match your queries.")
	fmt.Println("Benefits: No manual struct mapping, compile-time query validation, and excellent performance.")
}
