package check

import (
    "testing"
)

func TestGetModelSupportedModels(t *testing.T) {
    tests := []struct {
        name          string
        prompt        string
        providerName  string
        modelName     string
        key           string
        expectedModel string
    }{
        {"OpenAI GPT-4o Mini", "some prompt", "OpenAI", "gpt-4o-mini", "api-key", "gpt-4o-mini"},
        {"OpenAI GPT-3.5 Turbo", "some prompt", "OpenAI", "gpt-3.5-turbo", "api-key", "gpt-3.5-turbo"},
        {"GoogleAI Gemini 1.5 Flash", "prompt", "GoogleAI", "gemini-1.5-flash", "api-key", "gemini-1.5-flash"},
        {"Cohere Command-R", "prompt", "Cohere", "command-r", "api-key", "command-r"},
        {"Anthropic Claude-3 Sonnet", "prompt", "Anthropic", "claude-3-sonnet", "api-key", "claude-3-sonnet-20240229"}, // Updated expected model
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

