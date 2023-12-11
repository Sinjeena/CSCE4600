package builtins_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/Sinjeena/CSCE4600/Project2/builtins"
)

func Test_Pwd(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name       string
		args       args
		wantErr    error
		givenError error
		givenDir   string
	}{
		{
			name:       "pwd working",
			args:       args{args: []string{}},
			wantErr:    nil,
			givenError: nil,
			givenDir:   "./",
		},
		{
			name:       "pwd broken",
			args:       args{args: []string{}},
			wantErr:    nil,
			givenError: errors.New("broken"),
			givenDir:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out bytes.Buffer
			// testing
			if err := builtins.Pwd(tt.givenDir, tt.givenError, &out, tt.args.args...); tt.wantErr != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Pwd() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}
