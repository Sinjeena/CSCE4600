package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/user"
	"strings"

	"github.com/Sinjeena/CSCE4600/builtins"
)

func main() {
	exit := make(chan struct{}, 2) // buffer this so there's no deadlock.
	runLoop(os.Stdin, os.Stdout, os.Stderr, exit)
}

func runLoop(r io.Reader, w, errW io.Writer, exit chan struct{}) {
	var (
		input    string
		err      error
		readLoop = bufio.NewReader(r)
	)
	for {
		select {
		case <-exit:
			_, _ = fmt.Fprintln(w, "exiting gracefully...")
			return
		default:
			if err := printPrompt(w); err != nil {
				_, _ = fmt.Fprintln(errW, err)
				continue
			}
			if input, err = readLoop.ReadString('\n'); err != nil {
				_, _ = fmt.Fprintln(errW, err)
				continue
			}
			if err = handleInput(w, input, exit); err != nil {
				_, _ = fmt.Fprintln(errW, err)
			}
		}
	}
}

func printPrompt(w io.Writer) error {
	u, err := user.Current()
	if err != nil {
		return err
	}

	// Get the username without brackets.
	username := u.Username

	// Print the prompt with the current username.
	_, err = fmt.Fprintf(w, "[%v] $ ", username)
	return err
}

func handleInput(w io.Writer, input string, exit chan<- struct{}) error {
	// Remove trailing spaces.
	input = strings.TrimSpace(input)

	// Split the input separate the command name and the command arguments.
	args := strings.Split(input, " ")
	name, args := args[0], args[1:]

	// Check for built-in commands.
	// New builtin commands should be added here. Eventually this should be refactored to its own func.
	switch name {
	case "cd":
		return builtins.ChangeDirectory(args...)
	case "env":
		return builtins.EnvironmentVariables(w, args...)
	case "exit":
		exit <- struct{}{}
	case "echo":
		builtins.EchoCommand(args)
	case "pwd":
		builtins.PwdCommand()
	case "mkdir":
		if len(args) < 1 {
			fmt.Println("Directory name not provided.")
			// Handle the case where directory name is missing
			return nil
		}
		builtins.MkdirCommand(args[0])
	case "rmdir":
		if len(args) < 1 {
			fmt.Println("Directory name not provided.")
			// Handle the case where directory name is missing
			return nil
		}
		builtins.RmdirCommand(args[0]) // Pass the first argument as directory name
	case "touch":
		if len(args) < 1 {
			fmt.Println("File name not provided.")
			// Handle the case where file name is missing
			return nil
		}
		builtins.TouchCommand(args[0])
	}

	return executeCommand(name, args...)
}

func executeCommand(name string, arg ...string) error {
	// Otherwise prep the command
	cmd := exec.Command(name, arg...)

	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and return the error.
	return cmd.Run()
}
