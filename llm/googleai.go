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

func queryGoogleAI(prompt string, cfg *config.Config) (string, string, error) {
	justification := ""

	// Create a new context for API calls
	ctx := context.Background()

	// Create a new Google Generative AI client using the API key
	client, err := genai.NewClient(ctx, option.WithAPIKey(cfg.Project.LLM.ApiKey))
	if err != nil {
		log.Printf("Failed to create Google AI client: %v", err)
		return "", "", err
	}
	defer client.Close() // Ensure the client is closed when the function exits

	// Select and configure the generative model
	model := client.GenerativeModel(cfg.Project.LLM.Model)
	model.SetTemperature(float32(cfg.Project.LLM.Temperature)) // Set temperature
	model.SetCandidateCount(1)                                 // Set candidate count to 1
	cs := model.StartChat() // Start a new chat session
	model.ResponseMIMEType = "application/json"   

	// Add the initial user prompt to the chat history
	cs.History = []*genai.Content{
		{
			Parts: []genai.Part{
				genai.Text(prompt),
			},
			Role: "user",
		},
	}

	// Generate the initial content using the chat session
	resp, err := cs.SendMessage(ctx, genai.Text(prompt))
	if err != nil || len(resp.Candidates) == 0 {
		log.Printf("Completion error: err:%v \n", err)
		return "", "", fmt.Errorf("no response from Google AI: %v", err)
	}


	// Print the entire response object on log
	respJSON, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Println("Failed to marshal response:", err)
		return "", "", err
	}
	log.Printf("Full Google AI response: %s\n", string(respJSON))

	// Extract the content from the first candidate
	content := resp.Candidates[0].Content
	if content == nil || len(content.Parts) == 0 {
		log.Println("No content parts found in candidate")
		return "", "", fmt.Errorf("no content parts in response")
	}

	// Iterate over parts to find the text content
	var resultText string
	for _, part := range content.Parts {
		switch v := part.(type) {
		case genai.Text:
			resultText += string(v)
		default:
			log.Printf("Unhandled part type: %T\n", part)
		}
	}

	// Check if any text was extracted
	if resultText == "" {
		log.Println("No text content found in parts")
		return "", "", fmt.Errorf("no text content in response")
	}

	// If justification is required, send a follow-up message
	if cfg.Project.Configuration.CotJustification == "yes" {
		// Switch to text/plain MIME type for justification
		model.ResponseMIMEType = "text/plain"
		  
		justificationResp, err := cs.SendMessage(ctx, genai.Text(justification_query))
		if err != nil || len(justificationResp.Candidates) == 0 {
			log.Printf("Justification error: err:%v \n", err)
			return resultText, "", fmt.Errorf("no justification response from Google AI: %v", err)
		}

		// Print the entire response object on log
		respJ_JSON, err := json.MarshalIndent(justificationResp, "", "  ")
		if err != nil {
			log.Println("Failed to marshal justification response:", err)
			return "", "", err
		}
		log.Printf("Full Google AI justification response: %s\n", string(respJ_JSON))

		// Extract the justification content from the first candidate
		justificationContent := justificationResp.Candidates[0].Content
		if justificationContent == nil || len(justificationContent.Parts) == 0 {
			log.Println("No content parts found in justification response")
			return resultText, "", fmt.Errorf("no content parts in justification response")
		}

		// Iterate over parts to find the text content for justification
		for _, part := range justificationContent.Parts {
			switch v := part.(type) {
			case genai.Text:
				justification += string(v)
			default:
				log.Printf("Unhandled part type in justification: %T\n", part)
			}
		}

		// Check if any text was extracted for justification
		if justification == "" {
			log.Println("No text content found in justification parts")
			return resultText, "", fmt.Errorf("no text content in justification response")
		}
	}

	return resultText, justification, nil
}
