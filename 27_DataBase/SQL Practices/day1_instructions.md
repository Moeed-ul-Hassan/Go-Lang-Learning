# 📚 Day 1: The CRUD Architect

Welcome to your first day of Database Mastery! Today, we focus on the bread and butter of backend engineering: **CRUD**.

## 🎯 Objective
Build a robust Book Management system that can handle data safely and efficiently.

## 🛠️ Requirements
1.  **Prepared Statements**: Never use `fmt.Sprintf` to build queries. Use `?` placeholders.
2.  **Error Handling**: Every `err` must be checked.
3.  **Struct Mapping**: Use the `Book` struct to pass data around.

## 📝 Challenge Steps
1.  Open `day1_crud.go`.
2.  Complete the `TODO` items.
3.  Run your code using `go run day1_crud.go`.
4.  Verify that `library.db` is created and data is being persisted.

## 💡 Pro Tip
When using `rows.Scan()`, make sure the number of arguments matches the number of columns in your `SELECT` statement exactly!

---
*Ready to master the foundation? Let's code!*
