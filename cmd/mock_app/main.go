package main

import (
	app "gh.dev-tim/di-talk"
	"log"
)

func main() {
	// Initialize server with mock/staging dependencies
	server, cleanup, err := app.InitializeMockServer()
	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
		return
	}

	server.Start()
	defer cleanup()
}
