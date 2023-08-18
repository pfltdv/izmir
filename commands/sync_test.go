package commands

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSyncNoRoot(t *testing.T) {
	ConfigRoot = ""
	actual := new(bytes.Buffer)
	RootCmd.SetOut(actual)
	RootCmd.SetArgs([]string{"sync"})
	err := RootCmd.Execute()

	assert.ErrorContains(t, err, "Configuration file root does not exists")
}

func TestSyncRoot(t *testing.T) {
	ConfigRoot = ""
	actual := new(bytes.Buffer)
	RootCmd.SetOut(actual)
	RootCmd.SetArgs([]string{"sync", "-r", "/tmp"})
	err := RootCmd.Execute()

	assert.Nil(t, err)
	assert.Contains(t, actual.String(), "Platform Definition Folder is /tmp")
}
