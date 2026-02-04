package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Moeed-ul-Hassan/libraryapp/internal/data"
)

type Server struct {
	port  int
	books []data.Book // In-memory database for now
}

func NewServer(port int) *http.Server {
	s := &Server{
		port:  port,
		books: []data.Book{},
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      s.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
