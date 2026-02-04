-- Day 2: Relational Schema Design
-- We are expanding our Library system to include Members and Borrowing records.

-- 1. Members Table
CREATE TABLE IF NOT EXISTS members (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    joined_date DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- 2. Borrowings Table (The "Relationship" Table)
-- This table connects Books and Members.
CREATE TABLE IF NOT EXISTS borrowings (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    book_id INTEGER NOT NULL,
    member_id INTEGER NOT NULL,
    borrow_date DATE NOT NULL,
    return_date DATE,
    FOREIGN KEY (book_id) REFERENCES books(id),
    FOREIGN KEY (member_id) REFERENCES members(id)
);
