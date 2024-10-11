package llm

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"prismAId/config"
	"prismAId/cost"

	anthropic "github.com/anthropics/anthropic-sdk-go"
	option "github.com/anthropics/anthropic-sdk-go/option"
)

func queryAnthropic(prompt string, config *config.Config) (string, string, error) {
	justification := ""

	model := cost.GetModel(prompt, config)

	// Create a new Anthropic client
	client := anthropic.NewClient(
		option.WithAPIKey(config.Project.LLM.ApiKey), 
	)
	// Temperature?? float32(config.Project.LLM.Temperature)
	message, err := client.Messages.New(context.TODO(), anthropic.MessageNewParams{
		Model:     anthropic.F(model),
		MaxTokens: anthropic.F(int64(4096)),
		Messages: anthropic.F([]anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(prompt)),
		}),
	})
	if err != nil {
		return "", "", fmt.Errorf("no response from Anthropic: %v", err)
	}
	fmt.Printf("%+v\n", message.Content)

	// Print the entire response object on log
	respJSON, err := json.MarshalIndent(message.Content, "", "  ")
	if err != nil {
		log.Println("Failed to marshal response:", err)
		return "", "", err
	}
	log.Printf("Full Anthropic response: %s\n", string(respJSON))

	// Assuming the content response is what you typically use:
	if len(respJSON) == 0 || respJSON == nil {
		log.Println("No content found in response")
		return "", "", fmt.Errorf("no content in response")
	}

	answer := string(respJSON)

	/*if config.Project.Configuration.CotJustification == "yes" {
		// Continue the conversation to ask for justification within the same chat
		messages = append(messages, openai.ChatCompletionMessage{Role: openai.ChatMessageRoleUser, Content: justification_query})

		justificationParams := openai.ChatCompletionRequest{
			Model:       model,
			Messages:    messages, // Continue with the same conversation
			Temperature: float32(config.Project.LLM.Temperature),
		}

		justificationResp, err := client.CreateChatCompletion(context.Background(), justificationParams)
		if err != nil || len(justificationResp.Choices) != 1 {
			log.Printf("Justification error: err:%v len(choices):%v\n", err, len(justificationResp.Choices))
			return answer, "", fmt.Errorf("no justification response from OpenAI: %v", err)
		}
		
		// Assign the justification content
		if len(justificationResp.Choices) > 0 {
			justification = justificationResp.Choices[0].Message.Content
		} else {
			log.Println("No content found in justification response")
		}
	}*/

	return answer, justification, nil
}
