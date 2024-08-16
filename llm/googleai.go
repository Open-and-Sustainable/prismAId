package llm

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"prismAId/config"

	genai "github.com/google/generative-ai-go/genai"
	option "google.golang.org/api/option"
)

func queryGoogleAI(prompt string, cfg *config.Config) (string, error) {
	// Create a new context
	ctx := context.Background()

	// Create a new Google Generative AI client using the API key
	client, err := genai.NewClient(ctx, option.WithAPIKey(cfg.Project.LLM.ApiKey))
	if err != nil {
		log.Printf("Failed to create Google AI client: %v", err)
		return "", err
	}
	defer client.Close() // Ensure the client is closed when the function exits

	// Select and configure the generative model
	model := client.GenerativeModel(cfg.Project.LLM.Model)
	model.SetTemperature(float32(cfg.Project.LLM.Temperature)) // Set temperature
	model.SetCandidateCount(1)                                 // Set candidate count to 1
	model.ResponseMIMEType = "application/json"                // Set response format to JSON

	// Generate content using the configured model
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil || len(resp.Candidates) == 0 {
		log.Printf("Completion error: err:%v \n", err)
		return "", fmt.Errorf("no response from Google AI: %v", err)
	}

	// Print the entire response object on log
	respJSON, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Println("Failed to marshal response:", err)
		return "", err
	}
	log.Printf("Full Google AI response: %s\n", string(respJSON))

	// Check the content of the response
	if len(resp.Candidates) == 0 || resp.Candidates[0].Content == nil {
		log.Println("No content found in response")
		return "", fmt.Errorf("no content in response")
	}

	// Extract the content from the first candidate
	content := resp.Candidates[0].Content
	if content == nil || len(content.Parts) == 0 {
		log.Println("No content parts found in candidate")
		return "", fmt.Errorf("no content parts in response")
	}

	// Iterate over parts to find the text content
	var resultText string
	for _, part := range content.Parts {
		switch v := part.(type) {
		case genai.Text:
			// Since genai.Text is a type alias for string, we can directly concatenate it
			resultText += string(v)
		// Handle other cases if needed, like Blob, FunctionCall, etc.
		default:
			log.Printf("Unhandled part type: %T\n", part)
		}
	}

	// Check if any text was extracted
	if resultText == "" {
		log.Println("No text content found in parts")
		return "", fmt.Errorf("no text content in response")
	}

	return resultText, nil
}
