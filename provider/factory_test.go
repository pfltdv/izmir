package provider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAWSProvider(t *testing.T) {
	provider, err := GetCloudProvider(AWS)
	assert.NoError(t, err)
	assert.Equal(t, AWS, provider.Type())
}

func TestGetGCPProvider(t *testing.T) {
	provider, err := GetCloudProvider(GCP)
	assert.NoError(t, err)
	assert.Equal(t, GCP, provider.Type())
}

func TestInvalidGCPProvider(t *testing.T) {
	provider, err := GetCloudProvider("UNDEFINED")
	assert.Error(t, err)
	assert.Nil(t, provider)
}
