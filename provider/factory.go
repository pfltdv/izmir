package provider

import "fmt"

type CloudType string

const (
	AWS CloudType = "AWS"
	GCP CloudType = "GCP"
)

type ICloudProvider interface {
	Type() CloudType
}

func GetCloudProvider(cloudType CloudType) (ICloudProvider, error) {

	if cloudType == AWS {
		return newAWS()
	}
	if cloudType == GCP {
		return newGCP()
	}
	return nil, fmt.Errorf("Unsupported CloudType is passed")
}
