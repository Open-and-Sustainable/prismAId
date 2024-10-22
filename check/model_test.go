package check

import (
    "testing"
)

// TestGetModel tests the GetModel function for various providers and models.
func TestGetModel(t *testing.T) {
    tests := []struct {
        name          string
        prompt        string
        providerName  string
        modelName     string
        key           string
        expectedModel string
    }{
        {"OpenAI default", "some prompt", "OpenAI", "", "api-key", "gpt-4o-mini"},
        {"OpenAI specific model", "some prompt", "OpenAI", "gpt-3.5-turbo", "api-key", "gpt-3.5-turbo"},
        {"Unsupported provider", "some prompt", "SomeRandomAI", "", "api-key", ""},
        {"Unsupported model", "some prompt", "OpenAI", "unknown-model", "api-key", ""},

        // GoogleAI cases
        {"GoogleAI default high cost", "long prompt with many tokens", "GoogleAI", "", "api-key", "gemini-1.5-pro"},
        {"GoogleAI specific model", "prompt", "GoogleAI", "gemini-1.5-flash", "api-key", "gemini-1.5-flash"},
        {"GoogleAI unsupported model", "prompt", "GoogleAI", "unknown-model", "api-key", ""},

        // Cohere cases
        {"Cohere default", "prompt", "Cohere", "", "api-key", "command-r"},
        {"Cohere specific model", "prompt", "Cohere", "command", "api-key", "command"},
        {"Cohere unsupported model", "prompt", "Cohere", "unknown-model", "api-key", ""},

        // Anthropic cases
        {"Anthropic default", "prompt", "Anthropic", "", "api-key", "claude-3-haiku"},
        {"Anthropic specific model", "prompt", "Anthropic", "claude-3-sonnet", "api-key", "claude-3-sonnet"},
        {"Anthropic unsupported model", "prompt", "Anthropic", "unknown-model", "api-key", ""},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            model := GetModel(tt.prompt, tt.providerName, tt.modelName, tt.key)
            if model != tt.expectedModel {
                t.Errorf("GetModel(%q, %q, %q, %q) = %q; want %q", tt.prompt, tt.providerName, tt.modelName, tt.key, model, tt.expectedModel)
            }
        })
    }
}
