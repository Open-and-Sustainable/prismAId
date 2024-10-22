package check

import (
	"os"
	"strings"
	"testing"
)

func TestRunUserCheck(t *testing.T) {
	tests := []struct {
		name       string
		totalCost  string
		provider   string
		userInput  string
		wantErr    bool
		errMessage string
	}{
		{
			name:      "User agrees to proceed with GoogleAI",
			totalCost: "100",
			provider:  "GoogleAI",
			userInput: "y\n",
			wantErr:   false,
		},
		{
			name:       "User declines to proceed with OpenAI",
			totalCost:  "200",
			provider:   "OpenAI",
			userInput:  "n\n",
			wantErr:    true,
			errMessage: "operation aborted by the user",
		},
		{
			name:       "User provides invalid input with Anthropic",
			totalCost:  "300",
			provider:   "Anthropic",
			userInput:  "maybe\n",
			wantErr:    true,
			errMessage: "operation aborted by the user",
		},
		// Note: Simulating an input read error is non-trivial and may not be necessary for this test.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Save the original os.Stdin
			originalStdin := os.Stdin
			defer func() { os.Stdin = originalStdin }() // Ensure os.Stdin is restored after the test

			// Create a pipe to simulate os.Stdin
			r, w, err := os.Pipe()
			if err != nil {
				t.Fatalf("Failed to create pipe: %v", err)
			}

			// Write the test input to the writer end of the pipe
			_, err = w.WriteString(tt.userInput)
			if err != nil {
				t.Fatalf("Failed to write to pipe: %v", err)
			}
			// Close the writer to simulate end of input
			w.Close()

			// Assign the reader end of the pipe to os.Stdin
			os.Stdin = r

			// Call the function
			err = RunUserCheck(tt.totalCost, tt.provider)

			// Close the reader end after the function call
			r.Close()

			// Check for expected error
			if (err != nil) != tt.wantErr {
				t.Errorf("RunUserCheck() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil && tt.errMessage != "" && !strings.Contains(err.Error(), tt.errMessage) {
				t.Errorf("Expected error message to contain %q, got %q", tt.errMessage, err.Error())
			}
		})
	}
}
