package cost

import (
    "testing"

    "github.com/shopspring/decimal"
)

// MockTokenCounter implements the TokenCounter interface for testing
type MockTokenCounter struct {
    TokensPerPrompt map[string]int
}

func (m *MockTokenCounter) GetNumTokensFromPrompt(prompt string, provider string, model string, key string) int {
    if tokens, ok := m.TokensPerPrompt[prompt]; ok {
        return tokens
    }
    return 0 // Default to 0 if not specified
}

func TestComputeCosts(t *testing.T) {
    // Save the original tokenCounter and restore it after the test
    originalTokenCounter := tokenCounter
    defer func() { tokenCounter = originalTokenCounter }()

    // Set up the mock token counter
    mockCounter := &MockTokenCounter{
        TokensPerPrompt: map[string]int{
            "Test prompt one": 5,
            "Test prompt two": 6,
        },
    }

    // Replace the package-level tokenCounter with the mock
    tokenCounter = mockCounter

    prompts := []string{
        "Test prompt one",
        "Test prompt two",
    }

    provider := "OpenAI"
    model := "gpt-4"
    key := "test-api-key"

    // Expected total cost calculation
    numTokens1 := mockCounter.TokensPerPrompt["Test prompt one"]
    numTokens2 := mockCounter.TokensPerPrompt["Test prompt two"]

    cost1 := numCentsFromTokens(numTokens1, model)
    cost2 := numCentsFromTokens(numTokens2, model)
    expectedTotalCost := cost1.Add(cost2)

    // Call ComputeCosts
    totalCostStr := ComputeCosts(prompts, provider, model, key)

    totalCost, err := decimal.NewFromString(totalCostStr)
    if err != nil {
        t.Fatalf("Failed to parse total cost: %v", err)
    }

    if !totalCost.Equal(expectedTotalCost) {
        t.Errorf("Total cost mismatch. Expected: %s, Got: %s", expectedTotalCost.String(), totalCost.String())
    }
}
