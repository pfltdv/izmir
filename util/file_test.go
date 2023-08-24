package util

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirExists(t *testing.T) {
	localdir, _ := os.Getwd()
	assert.False(t, DirExists("/non_existing_directory"))
	assert.False(t, DirExists(localdir+"/util/file_test.go"))
	assert.True(t, DirExists(localdir))
}
