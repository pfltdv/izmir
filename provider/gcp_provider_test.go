package provider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGCP(t *testing.T) {
	provider, err := newGCP()
	assert.NotNil(t, provider)
	assert.NoError(t, err)
	assert.Equal(t, GCP, provider.Type())
}
