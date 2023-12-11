package builtins

import (
	"os"
	"testing"
	// import the package where MkdirCommand is defined
)

func TestMkdirCommand(t *testing.T) {
	testDir := "test_directory"

	// Clean up after the test completes
	defer func() {
		if err := os.RemoveAll(testDir); err != nil {
			t.Fatalf("Error cleaning up test directory: %s", err)
		}
	}()

	// Call the function
	if err := builtins.MkdirCommand(testDir); err != nil {
		t.Errorf("MkdirCommand() error = %v, want nil", err)
	}

	// Check if the directory exists
	if _, err := os.Stat(testDir); os.IsNotExist(err) {
		t.Errorf("MkdirCommand() did not create the directory")
	}
}
