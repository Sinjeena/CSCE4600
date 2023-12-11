package builtins

import (
	"errors"
	"fmt"
	"io"
)

var aliases = make(map[string][]string)

func HandleAlias(w io.Writer, args ...string) error {
	switch len(args) {
	case 0:
		for alias, command := range aliases {
			fmt.Fprintf(w, "%s expands to %s\n", alias, command)
		}
		if len(aliases) == 0 {
			fmt.Fprintf(w, "No aliases found\n")
		}
		return nil
	case 1:
		// print alias
		if val, ok := aliases[args[0]]; ok {
			fmt.Println(val)
			return nil
		}
		return errors.New("alias not found") // errors.New("alias chechers not found")
	default:
		// set alias
		aliases[args[0]] = args[1:]
		return nil
	}
}

func ReturnAliasMap() map[string][]string {
	return aliases
}
