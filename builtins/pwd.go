package builtins

import (
    "fmt"
    "os"
    "github.com/Sinjeena/Csce4600/Project2/builtins"
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
