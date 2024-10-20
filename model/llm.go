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

// QueryLLM sends a prompt to a specified LLM (Large Language Model) and retrieves the model's response,
// including justifications and summaries if applicable.
//
// Arguments:
// - prompt: A string containing the input prompt to be processed.
// - llm: A pointer to the LLM configuration being used.
// - options: A pointer to the review.Options configuration that specifies additional processing options.
//
// Returns:
// - Three strings representing the justification, summary, and full response from the model.
// - An error if the interaction with the LLM fails or if processing issues occur.
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

// NewLLM creates and returns a new LLM instance configured with the specified provider, model, and settings.
// This function initializes the LLM with API credentials and rate limits to interact with the respective provider.
//
// Arguments:
// - providerName: The name of the LLM service provider (e.g., "OpenAI", "Anthropic").
// - modelName: The specific name of the model to be used (e.g., "GPT-3", "Claude").
// - apiKey: A string containing the API key required for authenticating with the provider's service.
// - temperature: A float64 value representing the desired level of randomness in the model's responses.
// - tpm: The maximum number of tokens allowed per minute (Tokens Per Minute) for rate-limiting purposes.
// - rpm: The maximum number of requests allowed per minute (Requests Per Minute) for rate-limiting purposes.
// - id: A unique identifier for the created LLM instance.
//
// Returns:
// - A pointer to an LLM instance with the provided settings.
// - An error if the initialization fails.
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

