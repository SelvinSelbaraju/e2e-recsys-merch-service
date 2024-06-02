package server

import (
	triton "github.com/SelvinSelbaraju/e2e-recsys-merch-service/internal/client"
	"github.com/SelvinSelbaraju/e2e-recsys-merch-service/internal/feature_store"
)

import "net/http"

type Server interface {
	Start(address string, deps Dependencies)
	Stop()
}

type Dependencies struct {
	featureStore feature_store.FeatureStore
	// FIXME: This should be more generic
	inferenceServer triton.TritonClient
}

func CreateServer(address string) *http.Server {
	s := &http.Server{
		Addr: address,
	}
	return s
}
