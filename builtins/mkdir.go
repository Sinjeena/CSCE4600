package builtins

import (
    "fmt"
    "os"
    "github.com/Sinjeena/Csce4600/Project2/builtins"
)

// MkdirCommand represents the 'mkdir' command implementation
func MkdirCommand(directory string) {
    err := os.Mkdir(directory, 0755) // 0755 is the default permission for mkdir
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Directory", directory, "created successfully")
}
