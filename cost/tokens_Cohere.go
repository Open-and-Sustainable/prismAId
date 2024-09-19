package cost

import (
	"context"
	"log"

	"prismAId/config"

	cohere "github.com/cohere-ai/cohere-go/v2"
	cohereclient "github.com/cohere-ai/cohere-go/v2/client"
)

func numTokensFromPromptCohere(prompt string, modelName string, cfg *config.Config) (numTokens int) {
	// Create a new Cohere client
	client := cohereclient.NewClient(cohereclient.WithToken(cfg.Project.LLM.ApiKey))

	// Create the TokenizeRequest
	request := &cohere.TokenizeRequest{
		Text:  prompt, 
		Model: modelName,
	}

	// Call the Tokenize method
	response, err := client.Tokenize(context.Background(), request)
	if err != nil {
		log.Fatal(err)
	}

	// Return the number of tokens
	return len(response.Tokens)
}
