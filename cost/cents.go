package cost

import (
	"log"

	"github.com/shopspring/decimal"
)

// Define a map to hold the rates for each model
var modelRates = map[string]decimal.Decimal{ // dollar prices per input M token
	"gpt-4o-mini":             decimal.NewFromFloat(0.15).Div(decimal.NewFromInt(1000000)),
	"gpt-4o":                 decimal.NewFromFloat(5).Div(decimal.NewFromInt(1000000)),
	"gpt-4-turbo":            decimal.NewFromFloat(10).Div(decimal.NewFromInt(1000000)),
	"gpt-4":                  decimal.NewFromFloat(30).Div(decimal.NewFromInt(1000000)),
	"gpt-4-32k":              decimal.NewFromFloat(60).Div(decimal.NewFromInt(1000000)),
	"gpt-3.5-turbo":          decimal.NewFromFloat(0.5).Div(decimal.NewFromInt(1000000)),
	"gpt-3.5-turbo-instruct": decimal.NewFromFloat(1.5).Div(decimal.NewFromInt(1000000)),
	"gemini-1.5-flash":       decimal.NewFromFloat(0.15).Div(decimal.NewFromInt(1000000)), // the rate is halved if <= 128K input tokens, fixed below
	"gemini-1.5-pro":         decimal.NewFromFloat(7).Div(decimal.NewFromInt(1000000)),    // the rate is halved if <= 128K input tokens, fixed below
	"gemini-1.0-pro":         decimal.NewFromFloat(0.5).Div(decimal.NewFromInt(1000000)),
}

func numCentsFromTokens(numTokens int, model string) decimal.Decimal {
	rate, ok := modelRates[model]
	if !ok {
		rate = decimal.Zero
		log.Println("Cost estimation unavailable because model not found:", model)
	}
	// halve the rate if the number of tokens is less than 128K and using Google AI Gemini 1.5 models
	if numTokens <= 128000 && (model == "gemini-1.5-flash" || model == "gemini-1.5-pro") {
		rate = rate.Div(decimal.NewFromInt(2))
	}
	// Calculate the total cost in cents
	costInCents := decimal.NewFromInt(int64(numTokens)).Mul(rate)

	return costInCents
}
