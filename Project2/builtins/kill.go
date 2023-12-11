package builtins

import (
	"fmt"
	"io"
	"os/exec"
)

// Kill kills a process
func HandleKill(w io.Writer, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("format: kill [pid]")
	}
	cmd := exec.Command("kill", args[0])
	cmd.Stdout = w
	cmd.Stderr = w
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
