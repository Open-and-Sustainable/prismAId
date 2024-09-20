package cost

import (
	"context"
	"log"

	"prismAId/config"

	genai "github.com/google/generative-ai-go/genai"
	option "google.golang.org/api/option"
)

func numTokensFromPromptGoogleAI(prompt string, modelName string, cfg *config.Config) (numTokens int) {
	// Create a new context
	ctx := context.Background()
	// Create a new Google Generative AI client using the API key
	client, err := genai.NewClient(ctx, option.WithAPIKey(cfg.Project.LLM.ApiKey))
	if err != nil {
		log.Printf("Failed to create Google AI client: %v", err)
		return 0
	}
	defer client.Close() // Ensure the client is closed when the function exits
	// Select and configure the generative model
	model := client.GenerativeModel(modelName)
	tokResp, err := model.CountTokens(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}
	return int(tokResp.TotalTokens)
}
