package commands

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDropNoRoot(t *testing.T) {
	ConfigRoot = ""
	actual := new(bytes.Buffer)
	RootCmd.SetOut(actual)
	RootCmd.SetArgs([]string{"drop"})
	err := RootCmd.Execute()

	assert.ErrorContains(t, err, "Configuration file root does not exists")
}

func TestDropRoot(t *testing.T) {
	ConfigRoot = ""
	actual := new(bytes.Buffer)
	RootCmd.SetOut(actual)
	RootCmd.SetArgs([]string{"drop", "-r", "/tmp"})
	err := RootCmd.Execute()

	assert.Nil(t, err)
	assert.Contains(t, actual.String(), "Platform Definition Folder is /tmp")
}
