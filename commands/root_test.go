package commands

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootCommand(t *testing.T) {
	actual := new(bytes.Buffer)
	RootCmd.SetOut(actual)
	RootCmd.SetArgs([]string{})
	err := RootCmd.Execute()

	assert.Nil(t, err)
	assert.Contains(t, actual.String(), "completion")
	assert.Contains(t, actual.String(), "version")
	assert.Contains(t, actual.String(), "sync")
	assert.Contains(t, actual.String(), "diff")
	assert.Contains(t, actual.String(), "help")
}
