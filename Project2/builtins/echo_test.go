package builtins_test

import (
	"bytes"
	"testing"

	"github.com/Sinjeena/CSCE4600/Project2/builtins"
)

func Test_Echo(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name      string
		args      args
		wantAlias []string
		wantErr   error
	}{
		{
			name:    "echo empty",
			args:    args{args: []string{}},
			wantErr: nil,
		},
		{
			name:    "echo hello",
			args:    args{args: []string{"hello"}},
			wantErr: nil,
		},
		{
			name:    "echo hello world",
			args:    args{args: []string{"hello", "world"}},
			wantErr: nil,
		},
		{
			name:    "echo hello world !",
			args:    args{args: []string{"hello", "world", "!"}},
			wantErr: nil,
		},
		{
			name:    "echo hello world !",
			args:    args{args: []string{"hello", "world", "!", "!"}},
			wantErr: nil,
		},
		{
			name:    "echo hello world !",
			args:    args{args: []string{"hello", "world", "!", "!", "!"}},
			wantErr: nil,
		},
		{
			name:    "echo hello world !",
			args:    args{args: []string{"hello", "world", "!", "!", "!", "!"}},
			wantErr: nil,
		},
		{
			name:    "echo hello world !",
			args:    args{args: []string{"hello", "world", "!", "!", "!", "!", "!"}},
			wantErr: nil,
		},
		{
			name:    "echo hello world !",
			args:    args{args: []string{"hello", "world", "!", "!", "!", "!", "!", "!"}},
			wantErr: nil,
		},
		{
			name:    "echo hello world !",
			args:    args{args: []string{"hello", "world", "!", "!", "!", "!", "!", "!", "!"}},
			wantErr: nil,
		},
		{
			name:    "echo hello world !",
			args:    args{args: []string{"hello", "world", "!", "!", "!", "!", "!", "!", "!", "!"}},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out bytes.Buffer
			// testing
			if err := builtins.Echo(&out, tt.args.args...); tt.wantErr != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Echo() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}
