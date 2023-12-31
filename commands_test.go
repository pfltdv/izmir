package main

import (
	"bytes"
	"os"
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
	assert.Contains(t, actual.String(), "diff")
	assert.Contains(t, actual.String(), "drop")
	assert.Contains(t, actual.String(), "help")
	assert.Contains(t, actual.String(), "sync")
	assert.Contains(t, actual.String(), "version")
	assert.Contains(t, actual.String(), "-c, --config string")

}

func TestIzmirCommand(t *testing.T) {
	localdir, _ := os.Getwd()
	MainConfigRoot = ""
	ConfigRoot = ""
	actual := new(bytes.Buffer)
	RootCmd.SetOut(actual)
	RootCmd.SetArgs([]string{"diff", "-r", "/tmp", "-c", localdir + "/resources/config.yaml"})
	err := RootCmd.Execute()

	assert.Nil(t, err)
}

func TestVersion(t *testing.T) {
	actual := new(bytes.Buffer)
	RootCmd.SetOut(actual)
	RootCmd.SetArgs([]string{"version"})
	err := RootCmd.Execute()

	assert.Nil(t, err)
	assert.Equal(t, actual.String(), "Version:  \nBuild Number:  \nBuild Date:  \nCommit Hash:  \n")
}
