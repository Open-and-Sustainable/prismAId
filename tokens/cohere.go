package tokens

import (
	"context"
	"log"

	cohere "github.com/cohere-ai/cohere-go/v2"
	cohereclient "github.com/cohere-ai/cohere-go/v2/client"
)

func numTokensFromPromptCohere(prompt string, modelName string, key string) (numTokens int) {
	// Create a new Cohere client
	client := cohereclient.NewClient(cohereclient.WithToken(key))

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
