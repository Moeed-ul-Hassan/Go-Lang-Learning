package storage

import (
	"sync"

	"github.com/moeed/masteringapi/models"
)

// Storage acts as our "Database".
// Since we are not using MongoDB yet, we store data in the computer's RAM (memory).
type Storage struct {
	// tasks is a map where Key = Task ID (string) and Value = Task (struct).
	// A map is like a dictionary or hash table.
	tasks map[string]models.Task

	// mu is a Mutex (Mutual Exclusion).
	// Imagine 100 users trying to add a task at the EXACT same millisecond.
	// Without a Mutex, the map could get corrupted (two writes crashing into each other).
	// The Mutex acts like a "lock" on a bathroom door. Only one person (goroutine) can be inside at a time.
	// RWMutex is special: It allows MANY people to READ at once, but only ONE person to WRITE.
	mu sync.RWMutex
}

// NewStorage initializes our storage.
func NewStorage() *Storage {
	return &Storage{
		tasks: make(map[string]models.Task),
	}
}

// AddTask adds a new task to the store safely.
func (s *Storage) AddTask(t models.Task) {
	s.mu.Lock()         // LOCK the door! No one else can read or write right now.
	defer s.mu.Unlock() // UNLOCK the door when this function finishes (even if it crashes).

	s.tasks[t.ID] = t
}

// GetAllTasks returns a list of all tasks.
func (s *Storage) GetAllTasks() []models.Task {
	s.mu.RLock()         // READ LOCK! Many people can read, but no one can write.
	defer s.mu.RUnlock() // Unlock read access.

	var allTasks []models.Task
	for _, task := range s.tasks {
		allTasks = append(allTasks, task)
	}
	return allTasks
}

// GetTask returns a single task by ID.
func (s *Storage) GetTask(id string) (models.Task, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, found := s.tasks[id]
	return task, found
}

// UpdateTask updates an existing task.
func (s *Storage) UpdateTask(id string, updatedTask models.Task) bool {
	s.mu.Lock() // Write lock needed because we are changing data.
	defer s.mu.Unlock()

	_, found := s.tasks[id]
	if !found {
		return false
	}

	// Preserve the ID and CreatedAt from the original task if needed,
	// or just overwrite. Here we overwrite.
	s.tasks[id] = updatedTask
	return true
}

// DeleteTask removes a task.
func (s *Storage) DeleteTask(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, found := s.tasks[id]
	if !found {
		return false
	}

	delete(s.tasks, id)
	return true
}
