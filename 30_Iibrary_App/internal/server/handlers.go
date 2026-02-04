package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Moeed-ul-Hassan/libraryapp/internal/data"
)

func (s *Server) HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Production Ready Library System"))
}

func (s *Server) AddBookHandler(w http.ResponseWriter, r *http.Request) {
	var book data.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	book.ID = len(s.books) + 1
	s.books = append(s.books, book)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func (s *Server) ViewBooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s.books)
}

func (s *Server) IssueBookHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	for i, book := range s.books {
		if book.ID == id {
			if !book.Available {
				http.Error(w, "Book already issued", http.StatusBadRequest)
				return
			}
			s.books[i].Available = false
			w.Write([]byte("Book Issued Successfully"))
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

func (s *Server) ReturnBookHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	for i, book := range s.books {
		if book.ID == id {
			s.books[i].Available = true
			w.Write([]byte("Book Returned Successfully"))
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

func (s *Server) DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	for i, book := range s.books {
		if book.ID == id {
			s.books = append(s.books[:i], s.books[i+1:]...)
			w.Write([]byte("Book Deleted Successfully"))
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}
