package builtins_test

import (
	"bytes"
	"testing"

	"github.com/Sinjeena/CSCE4600/Project2/builtins"
)

func TestKill(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name:    "kill empty",
			args:    args{args: []string{}},
			wantErr: nil,
		},
		{
			name:    "specific kill doesn't exist",
			args:    args{args: []string{"6561658"}},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out bytes.Buffer
			// testing
			if err := builtins.HandleKill(&out, tt.args.args...); tt.wantErr != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("HandleKill() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}
