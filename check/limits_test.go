package check

import (
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
    return mtc.MockFunc(prompt, provider, model, key)
}

func TestRunInputLimitsCheck(t *testing.T) {
	mockCounter := MockTokenCounter{
        MockFunc: func(prompt string, provider string, model string, key string) int {
            switch model {
            case "gpt-3.5-turbo":
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
            model:    "gpt-3.5-turbo",
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
            prompt:   generateLongPrompt(20000), // Exceeds GPT35TurboMaxTokens
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

// generateLongPrompt generates a string with approximately numTokens words.
func generateLongPrompt(numTokens int) string {
    word := "word "
    var sb strings.Builder
    for i := 0; i < numTokens; i++ {
        sb.WriteString(word)
    }
    return sb.String()
}
