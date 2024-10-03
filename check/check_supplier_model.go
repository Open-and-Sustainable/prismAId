package check

import (
	"fmt"

	"prismAId/config"

	anthropic "github.com/anthropics/anthropic-sdk-go"
	openai "github.com/sashabaranov/go-openai"
)

const (
	// Supported Providers
	GoogleAI = "GoogleAI"
	OpenAI   = "OpenAI"
	Cohere   = "Cohere"
	Anthropic = "Anthropic"
)

var supportedProviders = map[string]bool{
	GoogleAI: true,
	OpenAI:   true,
	Cohere:   true,
	Anthropic: true,
}

var modelProviderMap = map[string]string{
    // OpenAI Models
    openai.GPT4oMini:      OpenAI,
    openai.GPT4o:           OpenAI,
    openai.GPT4Turbo:      OpenAI,
    openai.GPT3Dot5Turbo:    OpenAI,

    // GoogleAI Models
    "gemini-1.5-flash": GoogleAI,
    "gemini-1.5-pro":   GoogleAI,
    "gemini-1.0-pro":   GoogleAI,

    // Cohere Models
    "command-r-plus":   Cohere,
    "command-r":   	    Cohere,
	"command-light":    Cohere,
	"command":          Cohere,

	// Anthropic Models
	anthropic.ModelClaude_3_5_Sonnet_20240620:      Anthropic,
	anthropic.ModelClaude_3_Sonnet_20240229:      Anthropic,
	anthropic.ModelClaude_3_Opus_20240229:      Anthropic,
	anthropic.ModelClaude_3_Haiku_20240307:      Anthropic,
}

func IsModelInProvider(cfg *config.Config) error {
	model := cfg.Project.LLM.Model
	provider := cfg.Project.LLM.Provider

	// Check if the provider exists in the supported providers map
	// If the provider is not supported or cost mnimization is active, return appropriate response
	if _, exists := supportedProviders[provider]; !exists {
		return fmt.Errorf("provider %s is not supported", provider)
	} else if model == "" {
		// If the model is empty and the provider is valid, return nil
		return nil
	}

	// Check if the model exists in the map and if it belongs to the specified provider
	if modelProvider, exists := modelProviderMap[model]; exists {
		if modelProvider != provider {
			return fmt.Errorf("model %s does not belong to provider %s", model, provider)
		}
	} else {
		return fmt.Errorf("model %s is not recognized", model)
	}
	return nil
}

