# 🐬 MySQL: Basic to Advanced Guide

Welcome to the comprehensive guide for MySQL. This document covers everything from the absolute basics of relational databases to advanced optimization and Go integration.

---

## 🏗️ 1. The Foundation (Basics)

MySQL is a **Relational Database Management System (RDBMS)**. It stores data in tables with rows and columns, linked by unique keys.

### Data Definition Language (DDL)
Used to define the database structure.

- **CREATE**: Create a new database or table.
  ```sql
  CREATE DATABASE school_db;
  CREATE TABLE students (
      id INT AUTO_INCREMENT PRIMARY KEY,
      name VARCHAR(100) NOT NULL,
      age INT,
      email VARCHAR(100) UNIQUE
  );
  ```
- **ALTER**: Modify an existing table.
  ```sql
  ALTER TABLE students ADD COLUMN enrollment_date DATE;
  ```
- **DROP**: Delete a database or table.
  ```sql
  DROP TABLE students;
  ```

### Data Manipulation Language (DML)
Used to manage data within the tables.

- **INSERT**: Add new records.
  ```sql
  INSERT INTO students (name, age, email) VALUES ('Moeed', 22, 'moeed@example.com');
  ```
- **UPDATE**: Modify existing records.
  ```sql
  UPDATE students SET age = 23 WHERE name = 'Moeed';
  ```
- **DELETE**: Remove records.
  ```sql
  DELETE FROM students WHERE id = 1;
  ```

---

## 🔍 2. Data Query Language (DQL)

The most common part of SQL: retrieving data.

- **Basic Select**:
  ```sql
  SELECT name, email FROM students;
  ```
- **Filtering (WHERE)**:
  ```sql
  SELECT * FROM students WHERE age > 20 AND name LIKE 'M%';
  ```
- **Sorting & Limiting**:
  ```sql
  SELECT * FROM students ORDER BY age DESC LIMIT 5;
  ```

---

## 🔗 3. Intermediate: Relationships & Joins

Relational databases shine when connecting data across tables.

### Types of Joins
- **INNER JOIN**: Returns records that have matching values in both tables.
- **LEFT JOIN**: Returns all records from the left table, and matched records from the right.
- **RIGHT JOIN**: Returns all records from the right table, and matched records from the left.

```sql
SELECT students.name, courses.course_name
FROM students
INNER JOIN enrollments ON students.id = enrollments.student_id
INNER JOIN courses ON enrollments.course_id = courses.id;
```

### Aggregations
```sql
SELECT COUNT(*), AVG(age) FROM students;
SELECT department, COUNT(*) FROM teachers GROUP BY department HAVING COUNT(*) > 2;
```

---

## ⚡ 4. Advanced: Performance & Logic

### Indexes
Indexes speed up data retrieval but slow down writes. Use them on columns frequently used in `WHERE` clauses.
```sql
CREATE INDEX idx_student_email ON students(email);
```

### Transactions (ACID)
Ensure data integrity.
- **Atomicity**: All or nothing.
- **Consistency**: Valid state.
- **Isolation**: Concurrent transactions don't interfere.
- **Durability**: Saved permanently.

```sql
START TRANSACTION;
UPDATE accounts SET balance = balance - 100 WHERE id = 1;
UPDATE accounts SET balance = balance + 100 WHERE id = 2;
COMMIT; -- or ROLLBACK if something fails
```

### Normalization
The process of organizing data to reduce redundancy:
1. **1NF**: Atomic values, no repeating groups.
2. **2NF**: 1NF + all non-key attributes fully functional on the primary key.
3. **3NF**: 2NF + no transitive dependencies.

---

## 🐹 5. MySQL with Go

To use MySQL in Go, you need the driver:
`go get -u github.com/go-sql-driver/mysql`

### Connection Example
```go
import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // DSN: user:password@tcp(127.0.0.1:3306)/dbname
    db, err := sql.Open("mysql", "user:password@/school_db")
    if err != nil {
        panic(err)
    }
    defer db.Close()
}
```

> [!TIP]
> **Always use Prepared Statements** (`db.Prepare`) to prevent SQL Injection attacks. Never concatenate strings into your queries!

---
> [!IMPORTANT]
> **Database Mastery is about understanding how data flows, not just memorizing syntax. Keep practicing!**
