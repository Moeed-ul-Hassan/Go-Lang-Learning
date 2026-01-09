# Project: Contacts Book CLI

A comprehensive CLI application that combines almost all concepts learned in the first half of the course.

## Features
- **CRUD Operations**: Create, Read, Update, and Delete contacts.
- **Search**: Find contacts by name or email.
- **Data Integrity**: Ensure no duplicate emails or empty names.

## Concepts Combined
1. **Maps**: Used for fast lookup of contacts by ID or Email.
2. **Structs**: Defining the `Contact` entity.
3. **Methods**: Attaching logic like `Display()` or `Update()` to the `Contact` struct.
4. **Slices**: Listing and sorting contacts.
5. **Pointers**: Passing contacts to functions for modification.

## Project Structure
- `main.go`: Entry point and menu loop.
- `contact.go`: Struct definitions and methods.
- `manager.go`: Logic for managing the collection of contacts.

## Key Learnings
- Managing state in a CLI application.
- Choosing the right data structure for the job (Map vs Slice).
- Implementing search algorithms.
