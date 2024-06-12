package client

import (
	"context"
	triton "github.com/SelvinSelbaraju/e2e-recsys-merch-service/internal/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

type TritonClient struct {
	Client triton.GRPCInferenceServiceClient
}

func (tritonClient TritonClient) SendServerReadyRequest() {
	// Create context for our request with 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	serverReadyRequest := triton.ServerReadyRequest{}
	// Submit ServerReady request to server
	serverReadyResponse, err := tritonClient.Client.ServerReady(ctx, &serverReadyRequest)
	if err != nil {
		log.Fatalf("Couldn't get server ready: %v", err)
	}
	log.Printf("Triton Health - %v", serverReadyResponse)
}

func (tritonClient TritonClient) SendInferenceRequest(inferenceRequest *triton.ModelInferRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Printf("Sending input: %v", inferenceRequest)
	modelResponse, err := tritonClient.Client.ModelInfer(ctx, inferenceRequest)
	if err != nil {
		log.Fatalf("Received error: %v", err)
	}
	outputs := PostProcess(modelResponse)
	log.Printf("Received Outputs: %v", outputs)
}

func (tritonClient TritonClient) Init() {
	tritonClient.SendServerReadyRequest()
}

func NewTritonClient(conn *grpc.ClientConn) TritonClient {
	return TritonClient{
		Client: triton.NewGRPCInferenceServiceClient(conn),
	}
}
