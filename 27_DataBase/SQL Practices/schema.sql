-- This is an external SQL file
-- We will read this file in Go and execute it

CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users (username, email) VALUES ('admin', 'admin@example.com');
INSERT INTO users (username, email) VALUES ('moeed', 'moeed@example.com');
