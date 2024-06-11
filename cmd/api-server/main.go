package main

import (
	"github.com/SelvinSelbaraju/e2e-recsys-merch-service/internal/server"
	"log"
)

func main() {
	// Set up dependencies
	deps, err := server.CreateDependencies("test", "localhost:8001")
	if err != nil {
		log.Fatalf("dependency creation failed with error: %v", err)
	}
	deps.Init()

	// The server created will have handlers that use the dependencies
	server := server.CreateServer(":5001", deps)

	// Start the server
	log.Println("Starting HTTP server on port 5001")
	server.ListenAndServe()
}
