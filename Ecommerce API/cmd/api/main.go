package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Moeed-ul-Hassan/EcomAPI/config"
	"github.com/Moeed-ul-Hassan/EcomAPI/internal/router"
)

func main() {
	// Load config
	cfg := config.LoadConfig()

	// Setup router
	r := router.SetupRouter()

	// Start server
	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Server starting on %s in %s mode", addr, cfg.Env)
	
	err := http.ListenAndServe(addr, r)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
