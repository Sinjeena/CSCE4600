package builtins

import (
	"fmt"
	"os"
)

// RmdirCommand represents the 'rmdir' command implementation
func RmdirCommand(directory string) {
	err := os.Remove(directory)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Directory", directory, "removed successfully")
}
