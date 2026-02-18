package data

import "time"

type Book struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	ISBN      string    `json:"isbn"`
	Published time.Time `json:"published"`
	Available bool      `json:"available"`
}
