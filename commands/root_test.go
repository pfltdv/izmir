package commands

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootCommand(t *testing.T) {
	actual := new(bytes.Buffer)
	RootCmd.SetOut(actual)
	RootCmd.Execute()

	assert.Contains(t, actual.String(), "completion")
	assert.Contains(t, actual.String(), "version")
	assert.Contains(t, actual.String(), "help")
}
