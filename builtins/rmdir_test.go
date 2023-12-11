package builtins

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testWriter struct {
	io.Writer
	buf *bytes.Buffer
}

func (t *testWriter) Write(p []byte) (n int, err error) {
	return t.buf.Write(p)
}

func TestRmdirCommand(t *testing.T) {
	// Create a temporary directory for testing
	testDir := "test_directory"
	err := os.Mkdir(testDir, 0755)
	if err != nil {
		t.Fatalf("Error creating test directory: %s", err)
	}
	defer os.RemoveAll(testDir)

	// Backup the real os.Stdout for restoration later
	realStdout := os.Stdout

	// Create a buffer to capture stdout
	var buf bytes.Buffer
	capturedWriter := &testWriter{Writer: os.Stdout, buf: &buf}

	// Redirect stdout to the captured writer
	os.Stdout = capturedWriter

	RmdirCommand(testDir)

	// Restore os.Stdout to the original value
	os.Stdout = realStdout

	// Check if the test directory has been removed successfully
	_, err = os.Stat(testDir)
	assert.True(t, os.IsNotExist(err), "Directory should have been removed")

	expectedOutput := fmt.Sprintf("Directory %s removed successfully\n", testDir)
	assert.Equal(t, expectedOutput, buf.String(), "Output doesn't match expected message")
}
