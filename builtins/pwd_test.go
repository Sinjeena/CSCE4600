package builtins

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPwdCommand(t *testing.T) {
	var stdout bytes.Buffer
	os.Stdout = &stdout

	PwdCommand()

	// Get the expected output of the current working directory
	expected, err := os.Getwd()
	if err != nil {
		t.Fatalf("Error getting current directory: %s", err)
	}

	// Check if the printed output matches the expected current directory
	assert.Equal(t, expected+"\n", stdout.String(), "Output doesn't match expected directory")
}
