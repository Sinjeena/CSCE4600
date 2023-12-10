package builtins

import (
	"fmt"
	"os"
)

// PwdCommand represents the 'pwd' command implementation
func PwdCommand() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(pwd)
}
