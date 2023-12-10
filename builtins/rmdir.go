package builtins

import (
    "fmt"
    "os"
    "github.com/Sinjeena/Csce4600/Project2/builtins"
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
