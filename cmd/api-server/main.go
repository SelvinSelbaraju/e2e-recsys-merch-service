package main

import (
	"github.com/SelvinSelbaraju/e2e-recsys-merch-service/internal/client"
	"github.com/SelvinSelbaraju/e2e-recsys-merch-service/internal/connection"
	// "github.com/SelvinSelbaraju/e2e-recsys-merch-service/internal/server"
)

func main() {
	// server := server.CreateServer(":5001")
	// log.Println("Starting HTTP server on port 5001")
	// server.ListenAndServe()
	// Connect to gRPC server
	conn := connection.NewConnection("localhost:8001")
	defer conn.Close()

	// Create client from gRPC server connection
	client := client.NewTritonClient(conn)
	client.Init()
	client.SendInferenceRequest()
}
