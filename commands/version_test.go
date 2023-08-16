package commands

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	actual := new(bytes.Buffer)
	RootCmd.SetOut(actual)
	RootCmd.SetArgs([]string{"version"})
	err := RootCmd.Execute()

	assert.Nil(t, err)
	assert.Equal(t, actual.String(), "Version:  \nBuild Number:  \nBuild Date:  \nCommit Hash:  \n")
}
