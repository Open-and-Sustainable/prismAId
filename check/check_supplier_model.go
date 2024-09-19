package check

import (
	"fmt"

	"prismAId/config"
)

const (
	// Supported Providers
	GoogleAI = "GoogleAI"
	OpenAI   = "OpenAI"
	Cohere   = "Cohere"
)

var supportedProviders = map[string]bool{
	GoogleAI: true,
	OpenAI:   true,
	Cohere:   true,
}

var modelProviderMap = map[string]string{
    // OpenAI Models
    "gpt-4o-mini":      OpenAI,
    "gpt-4o":           OpenAI,
    "gpt-4-turbo":      OpenAI,
    "gpt-3.5-turbo":    OpenAI,

    // GoogleAI Models
    "gemini-1.5-flash": GoogleAI,
    "gemini-1.5-pro":   GoogleAI,
    "gemini-1.0-pro":   GoogleAI,

    // Cohere Models
    "command-r-plus":   Cohere,
    "command-r":   	    Cohere,
	"command-light":    Cohere,
	"command":          Cohere,
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

