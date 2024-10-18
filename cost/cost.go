package cost

import (
	"log"
	"prismAId/llm"

	"github.com/shopspring/decimal"
)

func ComputeCosts(prompts []string, llm *llm.LLM) string {
	// assess and report costs
	totalCost := decimal.NewFromInt(0)
	counter := 0
	for _, prompt := range prompts {
		counter++
		cost, err := assessPromptCost(prompt, llm)
		if err == nil {
			log.Println("File: ", counter, "Model: ", llm.Model, " Cost: ", cost)
			totalCost = totalCost.Add(cost)
		} else {
			log.Println("Error: ", err)
		}
	}
	return totalCost.String()
}

func assessPromptCost(prompt string, llm *llm.LLM) (decimal.Decimal, error) {
	numTokens := GetNumTokensFromPrompt(prompt, llm)
	numCents := numCentsFromTokens(numTokens, llm)
	return numCents, nil
}

func GetNumTokensFromPrompt(prompt string, llm *llm.LLM) int {
	var numTokens int
	switch llm.Provider {
	case "OpenAI":
		numTokens = numTokensFromPromptOpenAI(prompt, llm.Model, llm.APIKey)
	case "GoogleAI":
		numTokens = numTokensFromPromptGoogleAI(prompt, llm.Model, llm.APIKey)
	case "Cohere":
		numTokens = numTokensFromPromptCohere(prompt, llm.Model, llm.APIKey)
	case "Anthropic":
		numTokens = numTokensFromPromptOpenAI(prompt, "gpt-4o", llm.APIKey)
	default:
		log.Println("Unsupported LLM provider: ", llm.Provider)
		return 0
	}
	return numTokens
}
