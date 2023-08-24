package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFromFlag(t *testing.T) {
	localdir, _ := os.Getwd()
	os.Setenv(CONFIG_ENV_PARAM, localdir+"/../resources/broken.yaml")
	config, err := Load(localdir + "/../resources/config.yaml")
	assert.NoError(t, err)
	assert.NotNil(t, config)
}

func TestLoadFromEnvironment(t *testing.T) {
	localdir, _ := os.Getwd()
	os.Setenv(CONFIG_ENV_PARAM, localdir+"/../resources/config.yaml")
	config, err := Load("")
	assert.NoError(t, err)
	assert.NotNil(t, config)
}

func TestLoadNoConfig(t *testing.T) {
	os.Setenv(CONFIG_ENV_PARAM, "")
	config, err := Load("")
	assert.ErrorContains(t, err, "invalid configuration")
	assert.Nil(t, config)
}

func TestLoadBrokenConfig(t *testing.T) {
	localdir, _ := os.Getwd()
	config, err := Load(localdir + "/../resources/broken.yaml")
	assert.ErrorContains(t, err, "cannot unmarshal !!seq into config.IzmirConfig")
	assert.Nil(t, config)
}
