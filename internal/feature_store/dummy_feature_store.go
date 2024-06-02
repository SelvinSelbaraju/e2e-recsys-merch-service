package feature_store

import (
	"fmt"
	"log"
	"time"
)

var dummyUserData = map[string]UserFeatures{
	"1": {
		age:              21,
		clubMemberStatus: "ACTIVE",
	},
	"2": {
		age:              33,
		clubMemberStatus: "ACTIVE",
	},
	"3": {
		age:              52,
		clubMemberStatus: "ACTIVE",
	},
}

var dummyProductData = map[string]ProductFeatures{
	"1": {
		productTypeName:  "Bra",
		productGroupName: "Underwear",
		colourGroupName:  "Black",
		departmentName:   "Expressive Lingerie",
	},
	"2": {
		productTypeName:  "Sweater",
		productGroupName: "Garment Upper Body",
		colourGroupName:  "Pink",
		departmentName:   "Tops Knitwear DS",
	},
}

type DummyFeatureStore struct {
	url           string
	connectTimeMs int32
	latencyMs     int32
}

func NewDummyFeatureStore(url string, connectTimeMs int32, latencyMs int32) *DummyFeatureStore {
	return &DummyFeatureStore{
		url:           url,
		connectTimeMs: connectTimeMs,
		latencyMs:     latencyMs,
	}
}

func (f *DummyFeatureStore) Connect() {
	log.Printf("Mocking a connection for the dummy feature store at URL: %s", f.url)
	startTime := time.Now()
	time.Sleep(time.Duration(time.Duration(f.connectTimeMs) * time.Millisecond))
	duration := time.Since(startTime)
	log.Printf("Connected to Dummy Feature Store in %v ms", duration)
}

func (f *DummyFeatureStore) GetFeatures(featureType FeatureType, id string) (interface{}, error) {
	time.Sleep(time.Duration(f.latencyMs) * time.Millisecond)
	if featureType == "user" {
		features, exists := dummyUserData[id]
		if exists {
			return features, nil
		} else {
			return UserFeatures{}, nil
		}
	} else if featureType == "product" {
		features, exists := dummyProductData[id]
		if exists {
			return features, nil
		} else {
			return ProductFeatures{}, nil
		}
	} else {
		return nil, fmt.Errorf("features must be one of user or product, got %v", featureType)
	}
}
