package cost

import (
	"log"
	"strings"

	"github.com/shopspring/decimal"
)

// Define a map to hold the rates for each model
var modelRates = map[string]decimal.Decimal{ // dollar prices per M token
	"gpt-4-turbo":            decimal.NewFromFloat(10).Div(decimal.NewFromInt(1000000)),
	"gpt-4":                  decimal.NewFromFloat(30).Div(decimal.NewFromInt(1000000)),
	"gpt-4-32k":              decimal.NewFromFloat(60).Div(decimal.NewFromInt(1000000)),
	"gpt-3.5-turbo":          decimal.NewFromFloat(0.5).Div(decimal.NewFromInt(1000000)),
	"gpt-3.5-turbo-instruct": decimal.NewFromFloat(1.5).Div(decimal.NewFromInt(1000000)),
}

// NormalizeModelName attempts to map various external model names to known keys in the modelRates map.
func normalizeModelName(model string) string {
	model = strings.ToLower(model)
	if strings.Contains(model, "gpt-3.5") && strings.Contains(model, "instruct") {
		return "gpt-3.5-turbo-instruct" // Return a default or base model if specific one isn't known
	} else if strings.Contains(model, "gpt-3.5") {
		return "gpt-3.5-turbo"
	}
	if strings.Contains(model, "gpt-4") {
		if strings.Contains(model, "preview") || strings.Contains(model, "turbo") {
			return "gpt-4-turbo"
		} else {
			return "gpt-4" // Simplify all gpt-4 related queries to a base version
		}
	}
	return model
}

func numCentsFromTokens(numTokens int, model string) decimal.Decimal {
	normalizedModel := normalizeModelName(model)
	rate, ok := modelRates[normalizedModel]
	if !ok {
		rate = decimal.Zero
		log.Println("Cost estimation unavailable because model not found:", model)
	}

	// Calculate the total cost in cents
	costInCents := decimal.NewFromInt(int64(numTokens)).Mul(rate)

	return costInCents
}
