package util

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestConfig struct {
	Izmir map[string]interface{} `yaml:"izmir"`
	AWS   map[string]interface{} `yaml:"aws"`
	GCP   map[string]interface{} `yaml:"gcp"`
}

func TestDirExists(t *testing.T) {
	localdir, _ := os.Getwd()
	assert.False(t, DirExists("/non_existing_directory"))
	assert.False(t, DirExists(localdir+"/util/file_test.go"))
	assert.True(t, DirExists(localdir))
}

func TestUnmarshallNoFile(t *testing.T) {
	c := &TestConfig{}
	localdir, _ := os.Getwd()
	err := Unmarshall(localdir+"/../resources/non_existing_file", c)
	assert.ErrorContains(t, err, "no such file or directory")
}
func TestUnmarshallBrokenYaml(t *testing.T) {
	c := &TestConfig{}
	localdir, _ := os.Getwd()
	err := Unmarshall(localdir+"/../resources/broken.yaml", c)
	assert.ErrorContains(t, err, "cannot unmarshal !!seq into map[string]interface {}")
}

func TestUnmarshallYaml(t *testing.T) {
	c := &TestConfig{}
	localdir, _ := os.Getwd()
	err := Unmarshall(localdir+"/../resources/config.yaml", c)
	assert.NoError(t, err)
}
