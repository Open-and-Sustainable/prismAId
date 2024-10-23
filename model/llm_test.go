// mock_query_service.go

package model

import (
	"fmt"
	"github.com/Open-and-Sustainable/prismaid/review"
	"testing"
)

// MockQueryService simulates the QueryService for testing purposes
type MockQueryService struct{}

func (m MockQueryService) QueryLLM(prompt string, llm review.Model, options review.Options) (string, string, string, error) {
	// Simulate responses based on input for testing
	switch prompt {
	case "test prompt":
		return "justified based on prompt", "summary of prompt", "full response for prompt", nil
	default:
		return "", "", "", fmt.Errorf("no response for the given prompt")
	}
}

func TestSomeFunctionThatUsesQueryLLM(t *testing.T) {
	mockService := MockQueryService{}
	llm := review.Model{Provider: "MockProvider"}
	options := review.Options{}

	justification, summary, fullResponse, err := mockService.QueryLLM("test prompt", llm, options)
	if err != nil {
		t.Errorf("Failed during function execution: %v", err)
	}

	// Assert the expected outputs
	expectedJustification := "justified based on prompt"
	expectedSummary := "summary of prompt"
	expectedFullResponse := "full response for prompt"

	if justification != expectedJustification {
		t.Errorf("Expected justification %s, got %s instead", expectedJustification, justification)
	}
	if summary != expectedSummary {
		t.Errorf("Expected summary %s, got %s instead", expectedSummary, summary)
	}
	if fullResponse != expectedFullResponse {
		t.Errorf("Expected full response %s, got %s instead", expectedFullResponse, fullResponse)
	}
}