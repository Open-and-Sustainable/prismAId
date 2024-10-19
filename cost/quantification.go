package cost

import (
	"log"
	"prismAId/tokens"

	"github.com/shopspring/decimal"
)

func ComputeCosts(prompts []string, provider string, model string, key string) string {
	// assess and report costs
	totalCost := decimal.NewFromInt(0)
	counter := 0
	for _, prompt := range prompts {
		counter++
		cost, err := assessPromptCost(prompt, provider, model, key)
		if err == nil {
			log.Println("File: ", counter, "Model: ", model, " Cost: ", cost)
			totalCost = totalCost.Add(cost)
		} else {
			log.Println("Error: ", err)
		}
	}
	return totalCost.String()
}

func assessPromptCost(prompt string, provider string, model string, key string) (decimal.Decimal, error) {
	numTokens := tokens.GetNumTokensFromPrompt(prompt, provider, model, key)
	numCents := numCentsFromTokens(numTokens, model)
	return numCents, nil
}


