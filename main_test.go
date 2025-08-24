package main

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "quit immediately",
			input: "q\n",
		},
		{
			name:  "invalid word length then quit",
			input: "abc\nq\n",
		},
		{
			name:  "invalid word length then quit",
			input: "3\nq\n",
		},
		{
			name:  "valid word length then quit",
			input: "5\nq\nq\n",
		},
		{
			name:  "valid word then quit",
			input: "5\nag pb pb ly eb\nq\nq\n",
		},
		{
			name:  "invalid word then quit",
			input: "5\nag pb pb ly ec\nq\nq\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Redirect stdin
			oldStdin := os.Stdin
			r, w, _ := os.Pipe()
			os.Stdin = r

			// Redirect stdout to capture output
			oldStdout := os.Stdout
			rOut, wOut, _ := os.Pipe()
			os.Stdout = wOut

			// Write test input
			go func() {
				defer w.Close()
				fmt.Fprint(w, tt.input)
			}()

			// Run main function
			main()

			// Restore stdin/stdout
			os.Stdin = oldStdin
			os.Stdout = oldStdout
			wOut.Close()

			// Read output (optional)
			output, _ := io.ReadAll(rOut)
			if len(output) == 0 {
				t.Error("Expected some output from main")
			}
		})
	}
}
