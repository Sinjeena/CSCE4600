package builtins

import (
	"fmt"
	"io"
)

// pwd prints the current working directory
func Pwd(dir string, err error, w io.Writer, args ...string) error {
	wd, err := dir, err
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "%s\n", wd)
	return nil
}
