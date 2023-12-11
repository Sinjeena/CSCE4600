package builtins

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEchoCommand(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "single word",
			args:     []string{"echo", "Hello"},
			expected: "Hello\n",
		},
		{
			name:     "multiple words",
			args:     []string{"echo", "Welcome", "to", "Golang"},
			expected: "Welcome to Golang\n",
		},
		{
			name:     "empty args",
			args:     []string{"echo"},
			expected: "\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var capturedOutput bytes.Buffer
			oldStdout := capturedOutput // backup old stdout
			defer func() {
				capturedOutput = oldStdout // reset stdout
			}()

			EchoCommand(tt.args)
			result := capturedOutput.String()

			assert.Equal(t, tt.expected, result)
		})
	}
}
