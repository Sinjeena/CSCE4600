package builtins

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTouchCommand(t *testing.T) {
	testFileName := "testfile.txt"
	defer os.Remove(testFileName)

	TouchCommand(testFileName)

	// Check if the test file has been created successfully
	_, err := os.Stat(testFileName)
	assert.NoError(t, err, "File should have been created")

	// Check for the error when trying to create an already existing file
	TouchCommand(testFileName)
	assert.Error(t, err, "File already exists")
}
