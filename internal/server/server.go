package server

import (
	"encoding/json"
	"fmt"
	triton "github.com/SelvinSelbaraju/e2e-recsys-merch-service/internal/client"
	"github.com/SelvinSelbaraju/e2e-recsys-merch-service/internal/connection"
	"github.com/SelvinSelbaraju/e2e-recsys-merch-service/internal/feature_store"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"unicode"
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
	// dependencies.inferenceServer.SendInferenceRequest()
	// Get the features
	users, user_err := dependencies.featureStore.GetFeatures("user", []string{"1", "2"})
	products, product_err := dependencies.featureStore.GetFeatures("product", []string{"1", "2"})
	fmt.Printf("users: %v, products: %v", users, products)
	if user_err != nil || product_err != nil {
		log.Fatalf("feature fetching received prodct error: %v, user error: %v", product_err, user_err)
	}
	// Build the features into a request
	// First map any strings to float integers using the vocab map
	jsonFile, err := os.Open("vocab.json")
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var vocab map[string]map[string]float32
	json.Unmarshal([]byte(byteValue), &vocab)
	parsedUserFeatures := ParseFeatures(users, vocab)
	parsedProductFeatures := ParseFeatures(products, vocab)
	fmt.Printf("User: %v, Product: %v", parsedUserFeatures, parsedProductFeatures)
	bundledUsers := BundleFeatures(users)
	bundledProducts := BundleFeatures(products)
	fmt.Printf("users: %v, products: %v", bundledUsers, bundledProducts)
	// // Send a request with these features to the inference server
}

// Bundle features into map[string][]any from []any
func BundleFeatures(features []interface{}) map[string][]any {
	typeVal := reflect.ValueOf(features[0])
	result := make(map[string][]any)
	for i := 0; i < typeVal.NumField(); i++ {
		for _, bundle := range features {
			fieldName := reflect.TypeOf(bundle).Field(i).Name
			fieldValue := reflect.ValueOf(bundle).Field(i).Interface()
			result[fieldName] = append(result[fieldName], fieldValue)
		}
	}
	return result
}

func ParseFeatures(features []interface{}, vocab map[string]map[string]float32) []map[string]any {
	var parsedFeatures []map[string]any
	for _, bundle := range features {
		bundleMap := make(map[string]any)
		val := reflect.ValueOf(bundle)
		typ := reflect.TypeOf(bundle)
		for i := 0; i < val.NumField(); i++ {
			// Get the field name
			fieldName := typ.Field(i).Name
			fieldName = toSnakeCase(fieldName)

			// Get the field value
			fieldValue := val.Field(i).Interface()

			// Add the field to the map if the feature is present
			// If the category is not present use zero
			vocabMap, fieldExists := vocab[fieldName]
			if fieldExists {
				if stringField, ok := fieldValue.(string); ok {
					vocabVal, catExists := vocabMap[stringField]
					if catExists {
						bundleMap[fieldName] = vocabVal
					} else {
						bundleMap[fieldName] = 0
					}
				}
			} else {
				bundleMap[fieldName] = fieldValue
			}
		}
		parsedFeatures = append(parsedFeatures, bundleMap)
	}
	return parsedFeatures
}

func toSnakeCase(str string) string {
	var result strings.Builder
	for i, r := range str {
		// Check if the character is uppercase
		if unicode.IsUpper(r) {
			// Add an underscore before the uppercase letter if it's not the first character
			if i > 0 {
				result.WriteRune('_')
			}
			// Convert the character to lowercase
			result.WriteRune(unicode.ToLower(r))
		} else {
			// Otherwise, just add the character as it is
			result.WriteRune(r)
		}
	}
	return result.String()
}

func CreateServer(address string, dependencies Dependencies) *http.Server {
	http.HandleFunc("/", dependencies.Handler)
	s := &http.Server{
		Addr: address,
	}
	return s
}
