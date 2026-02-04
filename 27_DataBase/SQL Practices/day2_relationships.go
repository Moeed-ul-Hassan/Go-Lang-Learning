package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

/*
🌟 DAY 2 CHALLENGE: THE RELATIONSHIP EXPERT
Goal: Connect Books, Members, and Borrowing records.

Your Task:
1. Load the "day2_schema.sql" file to set up the new tables.
2. Implement a function to "Borrow a Book" (Insert into borrowings).
3. Implement a function to "Get All Borrowed Books" using an INNER JOIN.
   - Should return: Book Title, Member Name, and Borrow Date.
4. Implement a function to "Get Member Borrowing History" using a LEFT JOIN.
*/

type BorrowingRecord struct {
	BookTitle  string
	MemberName string
	BorrowDate string
}

func main() {
	db, err := sql.Open("sqlite", "library.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("--- Day 2: Relationship Mastery Started ---")

	// 1. TODO: Load and execute day2_schema.sql
	// 2. TODO: Implement BorrowBook(db, bookID, memberID, date)
	// 3. TODO: Implement GetBorrowedBooksReport(db)
	// 4. TODO: Implement GetMemberHistory(db, memberID)
}

// TODO: Define your functions here using Joins
// Example Join Query:
// SELECT books.title, members.name, borrowings.borrow_date
// FROM borrowings
// INNER JOIN books ON borrowings.book_id = books.id
// INNER JOIN members ON borrowings.member_id = members.id;
