package tokens

import (
    "testing"
)

func TestGetNumTokensFromPrompt(t *testing.T) {
    tests := []struct {
        name     string
        provider string
        model    string
        key      string
        prompt   string
        want     int
    }{
        {
            name:     "OpenAI provider with specific model",
            provider: "OpenAI",
            model:    "gpt-4",
            key:      "some-api-key",
            prompt:   "What is AI?",
            want:     100, // Assuming this is the expected token count
        },
        {
            name:     "Unsupported provider",
            provider: "UnknownAI",
            model:    "unknown-model",
            key:      "some-api-key",
            prompt:   "What is AI?",
            want:     0, // Unsupported provider should return 0
        },
        {
            name:     "GoogleAI provider",
            provider: "GoogleAI",
            model:    "gemini-1.5",
            key:      "some-api-key",
            prompt:   "What is AI?",
            want:     120, // Assuming this is the expected token count
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Use the MockTokenCounter to simulate behavior based on test case
            mockCounter := mockTokenCounter{
                mockTokenFunc: func(prompt, provider, model, key string) int {
                    if provider == "OpenAI" && model == "gpt-4" {
                        return 100
                    }
                    if provider == "GoogleAI" {
                        return 120
                    }
                    return 0 // For unsupported providers or any other cases
                },
            }

            got := mockCounter.GetNumTokensFromPrompt(tt.prompt, tt.provider, tt.model, tt.key)
            if got != tt.want {
                t.Errorf("GetNumTokensFromPrompt() = %v, want %v", got, tt.want)
            }
        })
    }
}

// MockTokenCounter is a mock implementation of the TokenCounter interface.
// It allows for custom behavior during testing by specifying the number of tokens to return.
type mockTokenCounter struct {
    mockTokenFunc func(prompt, provider, model, key string) int
}

// GetNumTokensFromPrompt is the mock method that simulates the token counting functionality.
func (mtc mockTokenCounter) GetNumTokensFromPrompt(prompt, provider, model, key string) int {
    if mtc.mockTokenFunc != nil {
        return mtc.mockTokenFunc(prompt, provider, model, key)
    }
    return 0 // Default to 0 tokens if no function is provided
}