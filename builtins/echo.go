package builtins

import (
    "fmt"
    "strings"
)

// EchoCommand represents the 'echo' command implementation
func EchoCommand(args []string) {
    fmt.Println(strings.Join(args[1:], " "))
}
