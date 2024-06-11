package server

import (
	"fmt"
	"net/http"

	triton "github.com/SelvinSelbaraju/e2e-recsys-merch-service/internal/client"
	"github.com/SelvinSelbaraju/e2e-recsys-merch-service/internal/connection"
	"github.com/SelvinSelbaraju/e2e-recsys-merch-service/internal/feature_store"
)

type Server interface {
	Start(address string, deps Dependencies)
	Stop()
}

type Dependencies struct {
	featureStore feature_store.FeatureStore
	// FIXME: This should be more generic
	inferenceServer triton.TritonClient
}

// FIXME: Should specify feature store type too
func CreateDependencies(featureStoreUrl string, inferenceServerUrl string) (Dependencies, error) {
	// Create grpc connection to inference server
	conn := connection.NewConnection(inferenceServerUrl)

	// Create client from gRPC server connection
	client := triton.NewTritonClient(conn)

	// Connect to the Feature Store
	// FIXME: This should be Feature Store agnostic
	featureStore, err := feature_store.CreateFeatureStore("dummy", "test")
	if err != nil {
		return Dependencies{}, fmt.Errorf("feature creation error: %v", err)
	}
	return Dependencies{
		featureStore:    featureStore,
		inferenceServer: client,
	}, nil
}

func (dependencies *Dependencies) Init() {
	// Test ping the connected inference server
	dependencies.inferenceServer.Init()

	// Connect to the Feature store
	dependencies.featureStore.Connect()
}

func (dependencies *Dependencies) Handler(w http.ResponseWriter, r *http.Request) {
	dependencies.inferenceServer.SendInferenceRequest()
}

func CreateServer(address string, dependencies Dependencies) *http.Server {
	http.HandleFunc("/", dependencies.Handler)
	s := &http.Server{
		Addr: address,
	}
	return s
}
