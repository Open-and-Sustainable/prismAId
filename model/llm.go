package model

import (
	"fmt"

	"prismAId/check"
	"prismAId/review"
)

const justification_query = "For each one of the keys and answers you provided, provide a justification for your answer as a chain of thought. In particular, I want a textual description of the few stages of the chin of thought that lead you to the answer you provided and the sentences in the text you analyzes that support your decision. If the value of a key was 'no' or empty '' because of lack of information on that topic in the text analyzed, explicitly report this reason. Please provide only th einformation requested, neither introductory nor concluding remarks."
const summary_query = "Summarize in very few sentences the text provided before for your review."

// LLM defines the structure for a large language model configuration
type LLM struct {
	Provider    string  // The provider of the LLM (e.g., "OpenAI", "Anthropic")
	Model       string  // The specific name of the model (e.g., "gpt-4", "claude-v1")
	APIKey      string  // API key for accessing the model
	Temperature float64 // Controls the randomness in the model's output
	TPM         int64   // Tokens per minute limit
	RPM         int64   // Requests per minute limit
	ID			string  // ID of the LLM, for ensemble purposes
}

func QueryLLM(prompt string, llm *LLM, options *review.Options) (string, string, string, error) {
	var queryFunc func(string, *LLM, *review.Options) (string, string, string, error)

	switch llm.Provider {
	case "OpenAI":
		queryFunc = queryOpenAI
	case "GoogleAI":
		queryFunc = queryGoogleAI
	case "Cohere":
		queryFunc = queryCohere
	case "Anthropic":
		queryFunc = queryAnthropic
	default:
		return "", "", "", fmt.Errorf("unsupported LLM provider: %s", llm.Provider)
	}

	return queryFunc(prompt, llm, options)
}

// Constructor-like function to create a cleaned and validated LLM object
func NewLLM(providerName, modelName, apiKey string, temperature float64, tpm, rpm int64, id string) (*LLM, error) {
	// get clean model name, an din the meanwhile check provider and model consistency
	modelName = check.GetModel("", providerName, modelName, apiKey)

	// Create and return the LLM object after validation
	return &LLM{
		Provider:    providerName,
		Model:       modelName,
		APIKey:      apiKey,
		Temperature: temperature,
		TPM:         tpm,
		RPM:         rpm,
		ID:			 id,
	}, nil
}

