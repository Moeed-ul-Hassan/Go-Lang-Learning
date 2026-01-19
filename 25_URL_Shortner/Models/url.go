package models

// Url represents the data structure for our URL Shortener.
// We use Capital letters (Exported) so other packages (like main and handlers) can access these fields.
// The `json:"..."` tags tell Go how to name these fields when converting to/from JSON.
type Url struct {
	ID       string `json:"id"`           // The short code (e.g., "a1b2")
	Original string `json:"original_url"` // The actual website (e.g., "https://google.com")
}
