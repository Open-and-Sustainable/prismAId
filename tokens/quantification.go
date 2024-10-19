package tokens

import (
	"log"
)


func GetNumTokensFromPrompt(prompt string, provider string, model string, key string) int {
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