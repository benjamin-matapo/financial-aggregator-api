package main

import (
	"log"
	"os"

	"financial-aggregator-api/backend/internal"
)

func main() {
	// Get port from environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create and start server
	server := internal.NewServer()

	log.Printf("Starting Financial Aggregator API on port %s", port)
	if err := server.Start(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
