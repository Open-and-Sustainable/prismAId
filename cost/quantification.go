package cost

import (
	"log"
	"prismAId/tokens"

	"github.com/shopspring/decimal"
)

// Declare a package-level TokenCounter variable
var tokenCounter tokens.TokenCounter = tokens.RealTokenCounter{}

// ComputeCosts processes a list of input prompts and calculates the total cost based on the specified 
// model and provider. The function uses predefined rates for each model and computes the cost by 
// iterating through each prompt.
//
// Arguments:
// - prompts: A slice of strings containing the input prompts to process.
// - provider: The name of the service provider (e.g., OpenAI).
// - model: The model being used to process the prompts.
// - key: An authentication key used for accessing the service.
//
// Returns:
// - A string containing the total cost information as a formatted output.
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
	numTokens := tokenCounter.GetNumTokensFromPrompt(prompt, provider, model, key)
	numCents := numCentsFromTokens(numTokens, model)
	return numCents, nil
}


