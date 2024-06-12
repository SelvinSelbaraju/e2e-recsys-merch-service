package feature_store

import (
	"fmt"
	"log"
	"time"
)

var dummyUserData = map[string]UserFeatures{
	"1": {
		Age:              21,
		ClubMemberStatus: "ACTIVE",
	},
	"2": {
		Age:              33,
		ClubMemberStatus: "ACTIVE",
	},
	"3": {
		Age:              52,
		ClubMemberStatus: "ACTIVE",
	},
}

var dummyProductData = map[string]ProductFeatures{
	"1": {
		ProductTypeName:  "Bra",
		ProductGroupName: "Underwear",
		ColourGroupName:  "Black",
		DepartmentName:   "Expressive Lingerie",
	},
	"2": {
		ProductTypeName:  "Sweater",
		ProductGroupName: "Garment Upper body",
		ColourGroupName:  "Pink",
		DepartmentName:   "Tops Knitwear DS",
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

func (f *DummyFeatureStore) GetFeatures(featureType FeatureType, ids []string) ([]interface{}, error) {
	time.Sleep(time.Duration(f.latencyMs) * time.Millisecond)
	var featureMaps []interface{}
	for _, id := range ids {
		if featureType == "user" {
			features, exists := dummyUserData[id]
			if exists {
				featureMaps = append(featureMaps, features)
			} else {
				featureMaps = append(featureMaps, UserFeatures{})
			}
		} else if featureType == "product" {
			features, exists := dummyProductData[id]
			if exists {
				featureMaps = append(featureMaps, features)
			} else {
				featureMaps = append(featureMaps, ProductFeatures{})
			}
		} else {
			return nil, fmt.Errorf("features must be one of user or product, got %v", featureType)
		}
	}
	return featureMaps, nil
}
