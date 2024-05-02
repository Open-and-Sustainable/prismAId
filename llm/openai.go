package llm

import (
	"context"
	"fmt"
	"os"
	"prismAId/cost"

	openai "github.com/sashabaranov/go-openai"
)

func QueryOpenAI(prompt string) (string, error) {

	//model := openai.GPT3Dot5Turbo // Adjust the model as necessary
	model := openai.GPT4TurboPreview

	// Create a new OpenAI client, loading the key from local file.
	apiKey := os.Getenv("OPENAI_API_KEY")
	client := openai.NewClient(apiKey)

	// Define your input data and create a prompt.
	messages := []openai.ChatCompletionMessage{{Role: openai.ChatMessageRoleUser, Content: prompt}}

	numTokens := cost.NumTokensFromMessages(messages, model)
	numCents := cost.NumCentsFromTokens(numTokens, model)

	// Print computed values
	fmt.Printf("Estimated Input Tokens: %d, Estimated Input Cost ($): %v\n", numTokens, numCents)

	// As if continuing
	check := cost.RunUserCheck()
	if check != nil {
		return "", nil
	}

	completionParams := openai.ChatCompletionRequest{
		Model:    model,
		Messages: messages,
		ResponseFormat: &openai.ChatCompletionResponseFormat{
			Type: openai.ChatCompletionResponseFormatTypeJSONObject, // Structured JSON response
		},
		//MaxTokens:   100, // to constraint the number of tokens in response
		Temperature: 0.5,
	}

	// Make the API call
	resp, err := client.CreateChatCompletion(context.Background(), completionParams)
	if err != nil || len(resp.Choices) != 1 {
		fmt.Printf("Completion error: err:%v len(choices):%v\n", err, len(resp.Choices))
		return "", fmt.Errorf("no response from OpenAI: %v", err)
	}

	// Print the entire response object
	/*respJSON, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		fmt.Println("Failed to marshal response:", err)
		return "", err
	}
	fmt.Printf("Full OpenAI response: %s\n", string(respJSON))*/

	// Assuming the content response is what you typically use:
	if len(resp.Choices) == 0 || resp.Choices[0].Message.Content == "" {
		fmt.Println("No content found in response")
		return "", fmt.Errorf("no content in response")
	}

	return resp.Choices[0].Message.Content, nil
}
