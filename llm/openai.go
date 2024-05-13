package llm

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"prismAId/config"
	"prismAId/cost"

	openai "github.com/sashabaranov/go-openai"
)

func QueryOpenAI(prompt string, config *config.Config) (string, error) {

	model := cost.GetModel(prompt, config)

	// Create a new OpenAI client
	client := openai.NewClient(config.Project.LLM.ApiKey)

	// Define your input data and create a prompt.
	messages := []openai.ChatCompletionMessage{{Role: openai.ChatMessageRoleUser, Content: prompt}}

	completionParams := openai.ChatCompletionRequest{
		Model:    model,
		Messages: messages,
		ResponseFormat: &openai.ChatCompletionResponseFormat{
			Type: openai.ChatCompletionResponseFormatTypeJSONObject,
		},
		Temperature: float32(config.Project.LLM.Temperature),
	}

	// Make the API call
	resp, err := client.CreateChatCompletion(context.Background(), completionParams)
	if err != nil || len(resp.Choices) != 1 {
		log.Printf("Completion error: err:%v len(choices):%v\n", err, len(resp.Choices))
		return "", fmt.Errorf("no response from OpenAI: %v", err)
	}

	// Print the entire response object on log
	respJSON, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Println("Failed to marshal response:", err)
		return "", err
	}
	log.Printf("Full OpenAI response: %s\n", string(respJSON))

	// Assuming the content response is what you typically use:
	if len(resp.Choices) == 0 || resp.Choices[0].Message.Content == "" {
		log.Println("No content found in response")
		return "", fmt.Errorf("no content in response")
	}

	return resp.Choices[0].Message.Content, nil
}
