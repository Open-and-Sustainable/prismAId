package llm

import (
	"context"
	"fmt"
	"log"
	"prismAId/config"
	"prismAId/cost"
	"strings"

	anthropic "github.com/anthropics/anthropic-sdk-go"
	option "github.com/anthropics/anthropic-sdk-go/option"
)

func queryAnthropic(prompt string, config *config.Config) (string, string, string, error) {
	justification := ""
	summary := ""

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
		return "", "", "", fmt.Errorf("no response from Anthropic: %v", err)
	}
	log.Printf("%+v\n", message.Content)

	// Accessing the content blocks directly
	textBlock := ""
	for _, block := range message.Content {
		// Check if the block is of type "text"
		if block.Type == "text" {
			textBlock = block.Text
		}
	}
	answer, err := extractSubstring(textBlock, "{", "}")
	if err != nil {
		return "", "", "", fmt.Errorf("no valid JSON review response from Anthropic: %v", err)
	}
	answer = "{\n"+answer+"\n}"

	if config.Project.Configuration.CotJustification == "yes" {
		justification, err = extractSubstringFrom(textBlock, "Explanation:")
		if err != nil {
			log.Printf("Justification error: err:%v \n", err)
			return answer, "", "", fmt.Errorf("no justification response from OpenAI: %v", err)
		} else if len(justification) == 0 {
			log.Println("No content found in justification response")
		}
	}

	log.Println(message.Content)
	log.Println(justification)

	return answer, justification, summary, nil
}

// Function to extract substring between first occurrences of two delimiters
func extractSubstring(s, startDelim, endDelim string) (string, error) {
	// Find the index of the first occurrence of the start delimiter
	startIndex := strings.Index(s, startDelim)
	if startIndex == -1 {
		return "", fmt.Errorf("start delimiter not found")
	}
	
	// Adjust the start index to skip over the start delimiter
	startIndex += len(startDelim)
	
	// Find the index of the first occurrence of the end delimiter after the start delimiter
	endIndex := strings.Index(s[startIndex:], endDelim)
	if endIndex == -1 {
		return "", fmt.Errorf("end delimiter not found")
	}
	
	// Adjust the endIndex relative to the original string
	endIndex += startIndex

	// Extract the substring between the two delimiters
	return s[startIndex:endIndex], nil
}

// Function to extract substring starting from the first occurrence of the start delimiter till the end of the string
func extractSubstringFrom(s, startDelim string) (string, error) {
	// Find the index of the first occurrence of the start delimiter
	startIndex := strings.Index(s, startDelim)
	if startIndex == -1 {
		return "", fmt.Errorf("start delimiter not found")
	}

	// Adjust the start index to skip over the start delimiter itself
	startIndex += len(startDelim)

	// Return the substring from the start delimiter to the end of the string
	return s[startIndex:], nil
}