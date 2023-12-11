package builtins_test

import (
	"bytes"
	"testing"

	"github.com/Sinjeena/CSCE4600/Project2/builtins"
)

func TestAlias(t *testing.T) {
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
			name:    "alias listing empty",
			args:    args{args: []string{}},
			wantErr: nil,
		},
		{
			name:    "specific alias print doesn't exist",
			args:    args{args: []string{"chechers"}},
			wantErr: nil,
		},
		{
			name:    "alias creation",
			args:    args{args: []string{"chechers", "echo", "hello"}},
			wantErr: nil,
		},
		{
			name:    "alias listing",
			args:    args{args: []string{}},
			wantErr: nil,
		},
		{
			name:    "specific alias printing",
			args:    args{args: []string{"chechers"}},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out bytes.Buffer
			// testing
			if err := builtins.HandleAlias(&out, tt.args.args...); tt.wantErr != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("HandleAlias() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
	if builtins.Unalias(nil, "") == nil {
		t.Errorf("Unalias() error = %v, wantErr %v", builtins.Unalias(nil, ""), nil)
	}
	if builtins.Unalias(nil, "chechers") != nil {
		t.Errorf("Unalias() error = %v, wantErr %v", builtins.Unalias(nil, "chechers"), nil)
	}
	if builtins.Unalias(nil, "chechers") == nil {
		t.Errorf("Unalias() error = %v, wantErr %v", builtins.Unalias(nil, "chechers"), nil)
	}
}
