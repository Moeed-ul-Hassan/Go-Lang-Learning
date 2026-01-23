package main

import (
	"fmt"
	"net/http"
)

// here we will learm MiddleWare in Go. Middleware is simply a function that wraps an http.Handler. It takes a handler as input, does some logic, and then calls that handler
func MiddleWare(next http.Handler) http.Handler {
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request))
	next.ServeHTTP(w, r)
	fmt.Println("Middleware")
	return http.HandlerFunc
}
