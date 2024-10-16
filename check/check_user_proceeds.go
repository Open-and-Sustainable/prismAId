package check

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"prismAId/config"
)

func RunUserCheck(totalCost string, cfg *config.Config) error {
	if  cfg.Project.LLM.Provider == "GoogleAI" {
		fmt.Println("Unless you are using a free tier with Google AI, the total cost (USD - $) to run this review is at least:", totalCost)
	} else if cfg.Project.LLM.Provider == "Anthropic" {
		fmt.Println("Anthropic tokenizer is not available, hence the estimation of the number of token is very imprecise.\nThe total cost (USD - $) to run this review should be at least:", totalCost)
	} else {
		fmt.Println("The total cost (USD - $) to run this review is at least:", totalCost)
	}
	fmt.Println("This value is an estimate of the total cost of input tokens only.")
	if cfg.Project.Configuration.CotJustification == "yes" {
		fmt.Println("Since you have chosen to include the CoT justifications of the answers provided, the total cost of inputs will be higher and depend on the cost of tokens stored in the chat.")
	}
	// Ask the user if they want to continue
	fmt.Print("Do you want to continue? (y/n): ")
	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read input: %v", err)
	}

	// Normalize and check response
	response = strings.TrimSpace(strings.ToLower(response))
	if response != "y" {
		return fmt.Errorf("operation aborted by the user")
	}

	return nil // No error, operation continues
}
