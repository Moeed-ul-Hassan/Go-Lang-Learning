package server

import (
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", s.HomeHandler)
	mux.HandleFunc("POST /add-book", s.AddBookHandler)
	mux.HandleFunc("GET /view-books", s.ViewBooksHandler)
	// Note: Generic handlers for now, can be improved with path values in Go 1.22
	mux.HandleFunc("/issue-book", s.IssueBookHandler)
	mux.HandleFunc("/return-book", s.ReturnBookHandler)
	mux.HandleFunc("/delete-book", s.DeleteBookHandler)

	return mux
}
