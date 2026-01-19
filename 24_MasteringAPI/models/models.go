package models

import "time"

// Model: A Model is just a blueprint for your data.
// Think of it like a "Class" in other languages or a database schema.
// It tells Go what fields a "Task" has (like ID, Name, Status) and what type they are.

// Task represents a to-do item in our application.
type Task struct {
	// ID is a unique identifier for the task.
	// `json:"id"` tells Go: "When converting this to JSON, name this field 'id' (lowercase)".
	ID string `json:"id"`

	// Title is the main content of the task.
	Title string `json:"title"`

	// Completed tracks if the task is done.
	Completed bool `json:"completed"`

	// CreatedAt stores when the task was created.
	CreatedAt time.Time `json:"created_at"`
}
