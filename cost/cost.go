package cost

import (
	"log"
	"prismAId/config"

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
	model := GetModel(prompt, config)
	numTokens := GetNumTokensFromPrompt(prompt, config)
	numCents := numCentsFromTokens(numTokens, model)
	return numCents, model, nil
}

func GetNumTokensFromPrompt(prompt string, cfg *config.Config) int {
	var numTokens int
	switch cfg.Project.LLM.Provider {
	case "OpenAI":
		model := GetModel(prompt, cfg)
		numTokens = numTokensFromPromptOpenAI(prompt, model)
	case "GoogleAI":
		model := GetModel(prompt, cfg)
		numTokens = numTokensFromPromptGoogleAI(prompt, model, cfg)
	case "Cohere":
		model := GetModel(prompt, cfg)
		numTokens = numTokensFromPromptCohere(prompt, model, cfg)
	case "Anthropic":
		numTokens = numTokensFromPromptOpenAI(prompt, "gpt-4o")
	default:
		log.Println("Unsupported LLM provider: ", cfg.Project.LLM.Provider)
		return 0
	}
	return numTokens
}
