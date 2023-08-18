package commands

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiffNoRoot(t *testing.T) {
	actual := new(bytes.Buffer)
	RootCmd.SetOut(actual)
	RootCmd.SetArgs([]string{"diff"})
	err := RootCmd.Execute()

	assert.ErrorContains(t, err, "Configuration file root does not exists")
}

func TestDiffRoot(t *testing.T) {
	actual := new(bytes.Buffer)
	RootCmd.SetOut(actual)
	RootCmd.SetArgs([]string{"diff", "-r", "/tmp"})
	err := RootCmd.Execute()

	assert.Nil(t, err)
	assert.Contains(t, actual.String(), "Platform Definition Folder is /tmp")
}
