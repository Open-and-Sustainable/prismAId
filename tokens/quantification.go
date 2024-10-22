package tokens

import (
	"log"
)

type TokenCounter interface {
    GetNumTokensFromPrompt(prompt string, provider string, model string, key string) int
}

// RealTokenCounter is the production implementation that uses the actual API.
type RealTokenCounter struct{}

// GetNumTokensFromPrompt determines the number of tokens in a given prompt based on the specified provider and model.
// This function uses provider-specific logic to count tokens using the respective provider's API or SDK.
//
// Arguments:
// - prompt: A string containing the input prompt to be analyzed.
// - provider: The name of the provider (e.g., "OpenAI", "Cohere", "GoogleAI").
// - model: The name of the model being used.
// - key: The API key required for authenticating with the provider.
//
// Returns:
// - An integer representing the number of tokens in the prompt.
func (rtc RealTokenCounter) GetNumTokensFromPrompt(prompt string, provider string, model string, key string) int {
    var numTokens int
    switch provider {
    case "OpenAI":
        numTokens = numTokensFromPromptOpenAI(prompt, model, key)
    case "GoogleAI":
        numTokens = numTokensFromPromptGoogleAI(prompt, model, key)
    case "Cohere":
        numTokens = numTokensFromPromptCohere(prompt, model, key)
    case "Anthropic":
        numTokens = numTokensFromPromptOpenAI(prompt, "gpt-4o", key)
    default:
        log.Println("Unsupported LLM provider: ", provider)
        return 0
    }
    return numTokens
}