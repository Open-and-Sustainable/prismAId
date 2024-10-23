package check

import (
    "fmt"
    "strings"
    "testing"

    anthropic "github.com/anthropics/anthropic-sdk-go"
    openai "github.com/sashabaranov/go-openai"
)

// MockTokenCounter is a mock implementation of the TokenCounter interface
type MockTokenCounter struct {
    MockFunc func(prompt string, provider string, model string, key string) int
}

func (mtc MockTokenCounter) GetNumTokensFromPrompt(prompt string, provider string, model string, key string) int {
    // Adjust token numbers to accurately trigger or not trigger errors
    switch model {
    case openai.GPT3Dot5Turbo:
        if strings.Contains(prompt, "20000") {  // Ensure this matches exactly what `generateLongPrompt(20000)` produces
            return 21000  // Must exceed the model's maximum tokens
        }
        return 1000  // Within normal limits
    case "unknown-model":
        return 0 // Simulate model not found
    case "gpt-4o-mini":
        return 130000 // Assumes this should exceed some set limit
    case "command":
        if strings.Contains(prompt, "Short prompt.") { // Specifically for the test case not expecting an error
            return 4000 // Below the limit of 4096
        }
        return 5000 // Above the limit to test failure scenarios in other cases
    case anthropic.ModelClaude_3_Haiku_20240307:
        return 250000 // Exceeds the limit for demonstration
    default:
        return 5000 // Default token count should not trigger errors unless specified
    }
}


func TestRunInputLimitsCheck(t *testing.T) {
	mockCounter := MockTokenCounter{
        MockFunc: func(prompt string, provider string, model string, key string) int {
            switch model {
            case openai.GPT3Dot5Turbo:
                return 1000 // Within limit
            case "unknown-model":
                return 0 // Simulate model not found
            case "gpt-4o-mini":
                return 130000 // Exceeds limit
            default:
                return 5000 // Default value
            }
        },
    }

    tests := []struct {
        name     string
        prompt   string
        provider string
        model    string
        key      string
        wantErr  bool
        errMsg   string
    }{
        {
            name:     "Valid input within limits - OpenAI GPT-3.5-Turbo",
            prompt:   "This is a short prompt.",
            provider: "OpenAI",
            model:    openai.GPT3Dot5Turbo,
            key:      "test-key",
            wantErr:  false,
        },
        {
            name:     "Model not found",
            prompt:   "This is a prompt.",
            provider: "OpenAI",
            model:    "unknown-model",
            key:      "test-key",
            wantErr:  true,
            errMsg:   "model 'unknown-model' not found",
        },
        {
            name:     "Prompt exceeds token limit - Anthropic Claude",
            prompt:   generateLongPrompt(250000), // Exceeds AnthropicMaxTokens
            provider: "Anthropic",
            model:    anthropic.ModelClaude_3_Haiku_20240307,
            key:      "test-key",
            wantErr:  true,
            errMsg:   "number of tokens in prompt",
        },
        {
            name:     "Valid input within limits - Cohere Command",
            prompt:   "Short prompt.",
            provider: "Cohere",
            model:    "command",
            key:      "test-key",
            wantErr:  false,
        },
        {
            name:     "Prompt exceeds token limit - OpenAI GPT-3.5-Turbo",
            prompt:   generateLongPrompt(20000),
            provider: "OpenAI",
            model:    openai.GPT3Dot5Turbo,
            key:      "test-key",
            wantErr:  true,
            errMsg:   "number of tokens in prompt",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := RunInputLimitsCheck(tt.prompt, tt.provider, tt.model, tt.key, mockCounter)
            if (err != nil) != tt.wantErr {
                t.Errorf("RunInputLimitsCheck() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if err != nil && !strings.Contains(err.Error(), tt.errMsg) {
                t.Errorf("Expected error message to contain %q, got %q", tt.errMsg, err.Error())
            }
        })
    }
}

// generateLongPrompt generates a string with approximately numTokens words plus the number.
func generateLongPrompt(numTokens int) string {
    word := "word "
    var sb strings.Builder
    for i := 0; i < numTokens; i++ {
        sb.WriteString(word)
    }
    // Append the exact number of tokens to the end of the prompt
    sb.WriteString(fmt.Sprintf(" %d", numTokens)) // Ensure this token count can be parsed or identified by the mock
    return sb.String()
}

