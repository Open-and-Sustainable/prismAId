package cost

import (
	"log"
	"prismAId/config"

	"github.com/sashabaranov/go-openai"
	"github.com/shopspring/decimal"
)

func ComputeCosts(prompts []string, config *config.Config) string {
	// assess and report costs
	totalCost := decimal.NewFromInt(0)
	counter := 0
	for _, prompt := range prompts {
		counter++
		cost, model, err := assessPromptCost(prompt, config)
		if err == nil {
			log.Println("File: ", counter, "Model: ", model, " Cost: ", cost)
			totalCost = totalCost.Add(cost)
		} else {
			log.Println("Error: ", err)
		}
	}
	return totalCost.String()
}

func assessPromptCost(prompt string, config *config.Config) (decimal.Decimal, string, error) {
	messages := []openai.ChatCompletionMessage{{Role: openai.ChatMessageRoleUser, Content: prompt}}
	model := GetModel(prompt, config)
	numTokens := numTokensFromMessages(messages, model)
	numCents := numCentsFromTokens(numTokens, model)
	return numCents, model, nil
}

func GetModel(prompt string, config *config.Config) string {
	//messages := []openai.ChatCompletionMessage{{Role: openai.ChatMessageRoleUser, Content: prompt}}
	model := openai.GPT4oMini
	if config.Project.LLM.Model == "" {
		/*numTokens := numTokensFromMessages(messages, model) // old code to minimize costs before GPT 4 Omni mini model availability
		if numTokens > 16385 {
			model = openai.GPT4o
		}*/
	} else if config.Project.LLM.Model == "gpt-3.5-turbo" {
		model = openai.GPT3Dot5Turbo
	} else if config.Project.LLM.Model == "gpt-4-turbo" {
		model = openai.GPT4Turbo
	} else if config.Project.LLM.Model == "gpt-4o" {
		model = openai.GPT4o
	} else if config.Project.LLM.Model == "gpt-4o-mini" {
		model = openai.GPT4oMini
	}
	return model
}

func GetNumTokensFromPrompt(prompt string, config *config.Config) int {
	messages := []openai.ChatCompletionMessage{{Role: openai.ChatMessageRoleUser, Content: prompt}}
	model := GetModel(prompt, config)
	return numTokensFromMessages(messages, model)
}
