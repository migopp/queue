package main

import (
	"github.com/migopp/queue/internal/api"
)

// Spawns server loop
func main() {
	// Create the server
	//
	// For now we hardcode for `localhost:8080` because
	// I felt like it, but we may want to programmatically
	// set this in the future
	server := api.Server{
		IP:   "localhost",
		Port: 8080,
	}
	if err := server.Run(); err != nil {
		// Maybe do something...
	}
}
