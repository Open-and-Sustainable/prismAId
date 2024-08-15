package cost

import (
	"context"
	"fmt"
	"log"
	"strings"

	"prismAId/config"

	genai "github.com/google/generative-ai-go/genai"
	tiktoken "github.com/pkoukk/tiktoken-go"
	openai "github.com/sashabaranov/go-openai"
	option "google.golang.org/api/option"
)

func numTokensFromPromptOpenAI(prompt string, model string) (numTokens int) {
	messages := []openai.ChatCompletionMessage{{Role: openai.ChatMessageRoleUser, Content: prompt}}
	tkm, err := tiktoken.EncodingForModel(model)
	if err != nil {
		err = fmt.Errorf("encoding for model: %v", err)
		log.Println(err)
		return 0 // Ensure consistent error handling by returning 0 tokens in case of error.
	}
	var tokensPerMessage, tokensPerName int
	switch model {
	case "gpt-3.5-turbo-0613",
		"gpt-3.5-turbo-16k-0613",
		"gpt-4-0314",
		"gpt-4-32k-0314",
		"gpt-4-0613",
		"gpt-4-32k-0613",
		"gpt-4o",
		"gpt-4oMini":
		tokensPerMessage = 3
		tokensPerName = 1
	case "gpt-3.5-turbo-0301":
		tokensPerMessage = 4
		tokensPerName = -1
	default:
		if strings.Contains(model, "gpt-3.5-turbo") {
			log.Println("warning: gpt-3.5-turbo may update over time. Returning num tokens assuming gpt-3.5-turbo-0613.")
			return numTokensFromPromptOpenAI(prompt, "gpt-3.5-turbo-0613")
		} else if strings.Contains(model, "gpt-4") {
			log.Println("warning: gpt-4 may update over time. Returning num tokens assuming computation as in gpt-4-0613, .")
			return numTokensFromPromptOpenAI(prompt, "gpt-4-0613")
		} else {
			err = fmt.Errorf("num_tokens_from_messages() is not implemented for model %s. See https://github.com/openai/openai-python/blob/main/chatml.md for information on how messages are converted to tokens", model)
			log.Println(err)
			return
		}
	}
	for _, message := range messages {
		numTokens += tokensPerMessage
		numTokens += len(tkm.Encode(message.Content, nil, nil))
		numTokens += len(tkm.Encode(message.Role, nil, nil))
		numTokens += len(tkm.Encode(message.Name, nil, nil))
		if message.Name != "" {
			numTokens += tokensPerName
		}
	}
	numTokens += 3 // replies are primed with <|start|>assistant<|message|>
	return numTokens
}

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
