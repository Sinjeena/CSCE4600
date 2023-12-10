package builtins

import (
	"fmt"
	"os"
)

// TouchCommand represents the 'touch' command implementation
func TouchCommand(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	fmt.Println("File", fileName, "created successfully")
}
