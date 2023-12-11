package builtins

import (
	"fmt"
	"os"
)

// MkdirCommand represents the 'mkdir' command implementation
func MkdirCommand(directory string) error {
	err := os.Mkdir(directory, 0755) // 0755 is the default permission for mkdir
	if err != nil {
		return fmt.Errorf("error creating directory: %v", err)
	}
	return nil
}
