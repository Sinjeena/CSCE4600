package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
	"testing/iotest"
	"time"

	"github.com/Sinjeena/CSCE4600/Project2/builtins"
	"github.com/stretchr/testify/require"
)

func Test_runLoop(t *testing.T) {
	t.Parallel()
	exitCmd := strings.NewReader("exit\n")
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name     string
		args     args
		wantW    string
		wantErrW string
	}{
		{
			name: "no error",
			args: args{
				r: exitCmd,
			},
		},
		//tests for builtins
		{
			name: "alias",
			args: args{
				r: strings.NewReader("alias\n"),
			},
			wantW: "EOF",
		},
		{
			name: "unalias",
			args: args{
				r: strings.NewReader("unalias\n"),
			},
		},
		{
			name: "pwd",
			args: args{
				r: strings.NewReader("pwd\n"),
			},
		},
		{
			name: "kill",
			args: args{
				r: strings.NewReader("kill\n"),
			},
		},
		{
			name: "cd",
			args: args{
				r: strings.NewReader("cd\n"),
			},
		},
		{
			name: "env",
			args: args{
				r: strings.NewReader("env\n"),
			},
		},
		{
			name: "echo",
			args: args{
				r: strings.NewReader("echo\n"),
			},
		},
		{
			name: "test",
			args: args{
				r: strings.NewReader("test\n"),
			},
		},
		{
			name: "read error should have no effect",
			args: args{
				r: iotest.ErrReader(io.EOF),
			},
			wantErrW: "EOF",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			w := &bytes.Buffer{}
			errW := &bytes.Buffer{}

			exit := make(chan struct{}, 2)
			// run the loop for 10ms
			go runLoop(tt.args.r, w, errW, exit)
			time.Sleep(10 * time.Millisecond)
			exit <- struct{}{}

			require.NotEmpty(t, w.String())
			if tt.wantErrW != "" {
				require.Contains(t, errW.String(), tt.wantErrW)
			}
		})
	}
}

func Test_RunAlias(t *testing.T) {
	var out bytes.Buffer
	if builtins.HandleAlias(&out, "chechers", "echo", "hello") != nil {
		return
	}
	if handleInput(&out, "chechers", nil) != nil {
		return
	}
	if (builtins.ReturnAliasMap()) == nil {
		return
	}
}

// func Test_InputHandler(t *testing.T) {
// 	var out bytes.Buffer
// 	if handleInput(&out, "alias", nil) != nil {
// 		return
// 	}
// 	if handleInput(&out, "unalias", nil) != nil {
// 		return
// 	}
// 	if handleInput(&out, "pwd", nil) != nil {
// 		return
// 	}
// 	if handleInput(&out, "kill", nil) != nil {
// 		return
// 	}
// 	if handleInput(&out, "cd", nil) != nil {
// 		return
// 	}
// 	if handleInput(&out, "env", nil) != nil {
// 		return
// 	}
// 	if executeCommand("testsdfasdf", "") != nil {
// 		return
// 	}
// }
