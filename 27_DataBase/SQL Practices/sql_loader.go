package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "modernc.org/sqlite"
)

/*
📖 HOW TO CONNECT SQL FILES WITH GO
In professional projects, we don't write long SQL queries as strings in Go.
Instead, we:
1. Write SQL in a .sql file (better syntax highlighting & formatting).
2. Read the file content in Go using os.ReadFile or ioutil.ReadFile.
3. Execute the resulting string using db.Exec().
*/

func main() {
	// 1. Connect to the database
	db, err := sql.Open("sqlite", "file_practice.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 2. Read the .sql file
	// Note: ioutil.ReadFile is deprecated in newer Go versions, use os.ReadFile
	content, err := ioutil.ReadFile("schema.sql")
	if err != nil {
		log.Fatal("Error reading SQL file:", err)
	}

	// 3. Convert bytes to string
	query := string(content)

	// 4. Execute the SQL
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal("Error executing SQL from file:", err)
	}

	fmt.Println("✅ Successfully executed SQL from schema.sql!")

	// Verify by querying the data
	rows, _ := db.Query("SELECT username FROM users")
	defer rows.Close()

	fmt.Println("Users in database:")
	for rows.Next() {
		var name string
		rows.Scan(&name)
		fmt.Printf("- %s\n", name)
	}
}
