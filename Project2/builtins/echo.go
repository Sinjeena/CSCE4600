package builtins

import (
	"fmt"
	"io"
)

// echo prints the arguments to the screen
func Echo(w io.Writer, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("no arguments provided")
	}
	for _, arg := range args {
		fmt.Fprintf(w, "%s ", arg)
	}
	fmt.Fprintf(w, "\n")
	return nil
}
