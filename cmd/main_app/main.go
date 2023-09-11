package main

import (
	app "gh.dev-tim/di-talk"
	"log"
)

func main() {
	// Initialize server
	server, cleanup, err := app.InitializeServer()
	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
		return
	}

	server.Start()
	defer cleanup()
}
