package provider

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAWS(t *testing.T) {
	provider, err := newAWS()
	assert.NotNil(t, provider)
	assert.NoError(t, err)
	assert.Equal(t, AWS, provider.Type())
}
