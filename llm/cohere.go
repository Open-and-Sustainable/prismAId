package llm

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"prismAId/config"
	"prismAId/cost"

	cohere "github.com/cohere-ai/cohere-go/v2"
	cohereclient "github.com/cohere-ai/cohere-go/v2/client"
	uuid "github.com/google/uuid"
)

func queryCohere(prompt string, config *config.Config) (string, string, error) {
	justification := ""

	model := cost.GetModel(prompt, config)

	chatID := uuid.New().String()

	// Create a new Cohere client
	client := cohereclient.NewClient(cohereclient.WithToken(config.Project.LLM.ApiKey))

	// Define your input data and create a prompt
	chatRequest := &cohere.ChatRequest{
		Message:        prompt,
		Model:          &model,
		ConversationId: &chatID,
		Temperature:    &config.Project.LLM.Temperature,
	}

	// Make the API call
	response, err := client.Chat(context.TODO(), chatRequest)
	if err != nil {
		log.Printf("Completion error: err:%v len(generations):%v\n", err, len(response.Text))
		return "", "", fmt.Errorf("no response from Cohere: %v", err)
	}

	// Print the entire response object on log
	respJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Println("Failed to marshal response:", err)
		return "", "", err
	}
	log.Printf("Full Cohere response: %s\n", string(respJSON))

	if len(response.Text) == 0 || response.Text == "" {
		log.Println("No content found in response")
		return "", "", fmt.Errorf("no content in response")
	}

	answer := response.Text

	if config.Project.Configuration.CotJustification == "yes" {
		// Continue the conversation to ask for justification within the same chat
		justificationRequest := &cohere.ChatRequest{
			Message:        justification_query,             // The query for justification
			Model:          &model,                          // Same model
			ConversationId: &chatID,                         // Continue with the same chat ID
			Temperature:    &config.Project.LLM.Temperature, // Same temperature
		}

		// Make the API call to ask for justification
		justificationResponse, err := client.Chat(context.TODO(), justificationRequest)
		if err != nil || justificationResponse.Text == "" {
			log.Printf("Justification error: err:%v len(text):%v\n", err, len(justificationResponse.Text))
			return answer, "", fmt.Errorf("no justification response from Cohere: %v", err)
		}

		// Assign the justification content from the response
		justification = justificationResponse.Text
	}

	return answer, justification, nil
}
