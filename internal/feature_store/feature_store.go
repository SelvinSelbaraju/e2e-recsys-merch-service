package feature_store

type FeatureStore interface {
	Connect()
	GetFeatures(featureType FeatureType, id string) (interface{}, error)
}

type FeatureType string

const (
	UserFeatureType    FeatureType = "user"
	ProductFeatureType FeatureType = "product"
)

type UserFeatures struct {
	clubMemberStatus string
	age              int8
}

type ProductFeatures struct {
	productTypeName  string
	productGroupName string
	colourGroupName  string
	departmentName   string
}
