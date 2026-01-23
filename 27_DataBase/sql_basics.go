package main

import (
	"database/sql" //Importing SQL
	"fmt"
	"log"
	"time"

	_ "modernc.org/sqlite" // Import the SQLite driver
)

/*
All code in this file is AI written
Database Mastery: SQL Fundamentals
This file demonstrates the core concepts of working with SQL databases in Go.
*/

func sql_basics() {
	// 1. Establishing connections: Use sql.Open() with a driver-specific DSN
	// For SQLite, the DSN is just the filename.
	db, err := sql.Open("sqlite", "learning.db") // learning.db is the name of the database. also this a gateway for database.
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	defer db.Close() //If there is an error while starting database it will close the database.

	// 2. Connection pooling: Configure pool size
	// The database/sql package manages this automatically.
	db.SetMaxOpenConns(10)           // Maximum number of open connections to the database.
	db.SetMaxIdleConns(5)            // Maximum number of connections in the idle connection pool.
	db.SetConnMaxLifetime(time.Hour) // Maximum amount of time a connection may be reused.

	fmt.Println("Successfully connected to SQLite database!")

	// Create a table for demonstration
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		age INTEGER
	)`)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}

	// 3. Executing queries: Use db.Exec() for INSERT, UPDATE, and DELETE
	res, err := db.Exec("INSERT INTO students (name, age) VALUES (?, ?)", "Moeed", 22)
	if err != nil {
		log.Fatal("Error inserting data:", err)
	}
	lastID, _ := res.LastInsertId()
	fmt.Printf("Inserted student with ID: %d\n", lastID)

	// 4. Handling results: Use rows.Scan() to read data from query results
	// Use db.Query() for SELECT statements
	rows, err := db.Query("SELECT id, name, age FROM students")
	if err != nil {
		log.Fatal("Error querying data:", err)
	}
	defer rows.Close()

	fmt.Println("Current Students:")
	for rows.Next() {
		var id int
		var name string
		var age int
		// Scan copies the columns in the current row into the values pointed at by dest.
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal("Error scanning row:", err)
		}
		fmt.Printf("- ID: %d, Name: %s, Age: %d\n", id, name, age)
	}

	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	// 5. Prepared statements: Use db.Prepare() to avoid SQL injection
	stmt, err := db.Prepare("INSERT INTO students (name, age) VALUES (?, ?)")
	if err != nil {
		log.Fatal("Error preparing statement:", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec("Alice", 20)
	if err != nil {
		log.Fatal("Error executing prepared statement:", err)
	}
	fmt.Println("Inserted Alice using a prepared statement.")

	// 6. Transactions: Use db.Begin() to start a transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatal("Error starting transaction:", err)
	}

	// Inside a transaction, use tx.Exec() or tx.Query()
	_, err = tx.Exec("INSERT INTO students (name, age) VALUES (?, ?)", "Bob", 25)
	if err != nil {
		// Rollback if there's an error
		tx.Rollback()
		log.Fatal("Transaction failed, rolled back:", err)
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal("Error committing transaction:", err)
	}
	fmt.Println("Transaction committed successfully (Inserted Bob).")
}
