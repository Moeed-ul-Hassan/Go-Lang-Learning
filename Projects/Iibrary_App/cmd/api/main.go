package main

import (
	"fmt"

	"github.com/Moeed-ul-Hassan/libraryapp/internal/server"
)

func main() {

	server := server.NewServer(8080)

	fmt.Println("Starting Production Ready Library System on :8080")
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
