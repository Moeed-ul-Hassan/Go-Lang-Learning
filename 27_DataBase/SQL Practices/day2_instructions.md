# 🔗 Day 2: The Relationship Expert

Today we move from single tables to **Relational Data**. In the real world, data is almost always connected.

## 🎯 Objective
Master the art of connecting tables using **Foreign Keys** and retrieving data using **Joins**.

## 🛠️ Requirements
1.  **Foreign Keys**: Ensure your `borrowings` table correctly references `books` and `members`.
2.  **Joins**: 
    - Use `INNER JOIN` when you only want records that exist in both tables.
    - Use `LEFT JOIN` when you want all records from the first table, even if they don't have a match in the second.
3.  **Data Integrity**: Try to insert a borrowing for a `book_id` that doesn't exist. What happens? (Hint: Foreign Key constraint should stop you!).

## 📝 Challenge Steps
1.  Open `day2_relationships.go`.
2.  Load the `day2_schema.sql` file using the technique we learned today (Reading files in Go).
3.  Implement the `TODO` functions.
4.  Run your code and verify the output.

## 💡 Pro Tip
When writing Joins, always use table aliases (e.g., `FROM books AS b`) to keep your queries short and readable!

---
*Relationships are the heart of SQL. Master them, and you master the database!*
