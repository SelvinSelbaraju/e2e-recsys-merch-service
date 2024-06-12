package feature_store

import (
	"fmt"
)

type FeatureStore interface {
	Connect()
	GetFeatures(featureType FeatureType, ids []string) ([]interface{}, error)
}

type FeatureType string

const (
	UserFeatureType    FeatureType = "user"
	ProductFeatureType FeatureType = "product"
)

type UserFeatures struct {
	ClubMemberStatus string
	Age              int8
}

type ProductFeatures struct {
	ProductTypeName  string
	ProductGroupName string
	ColourGroupName  string
	DepartmentName   string
}

func CreateFeatureStore(featureStoreType string, url string) (FeatureStore, error) {
	if featureStoreType == "dummy" {
		return &DummyFeatureStore{
			url:           url,
			connectTimeMs: 100,
			latencyMs:     10,
		}, nil
	} else {
		return nil, fmt.Errorf("featureStoreType must be one of dummy, got %v", featureStoreType)
	}
}
