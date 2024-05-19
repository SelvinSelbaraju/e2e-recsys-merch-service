package main

import (
	"context"
	"fmt"
	triton "github.com/SelvinSelbaraju/e2e-recsys-merch-service/internal/proto"
	// "github.com/SelvinSelbaraju/e2e-recsys-merch-service/internal/server"
	"google.golang.org/grpc"
	"log"
	"time"
)

func ServerReadyRequest(client triton.GRPCInferenceServiceClient) *triton.ServerReadyResponse {
	// Create context for our request with 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	serverReadyRequest := triton.ServerReadyRequest{}
	// Submit ServerReady request to server
	serverReadyResponse, err := client.ServerReady(ctx, &serverReadyRequest)
	if err != nil {
		log.Fatalf("Couldn't get server ready: %v", err)
	}
	return serverReadyResponse
}

func main() {
	// server := server.CreateServer(":5001")
	// log.Println("Starting HTTP server on port 5001")
	// server.ListenAndServe()
	// Connect to gRPC server
	const URL = "localhost:8001"
	conn, err := grpc.Dial(URL, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't connect to endpoint %s: %v", URL, err)
	}
	defer conn.Close()

	// Create client from gRPC server connection
	client := triton.NewGRPCInferenceServiceClient(conn)

	serverReadyResponse := ServerReadyRequest(client)
	fmt.Printf("Triton Health - Ready: %v\n", serverReadyResponse.Ready)
}
