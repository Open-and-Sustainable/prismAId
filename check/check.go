package check

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"prismAId/config"
	"prismAId/cost"
)

const (
    // OpenAI Models
    GPT4MiniMaxTokens      = 128000
    GPT4MaxTokens          = 128000
    GPT4TurboMaxTokens     = 128000
    GPT35TurboMaxTokens    = 16385
    // GoogleAI Models
    Gemini15FlashMaxTokens = 1048576
    Gemini15ProMaxTokens   = 2097152
    Gemini10ProMaxTokens   = 32760
)

var ModelMaxTokens = map[string]int{
    "gpt-4o-mini":      GPT4MiniMaxTokens,
    "gpt-4o":           GPT4MaxTokens,
    "gpt-4-turbo":      GPT4TurboMaxTokens,
    "gpt-3.5-turbo":    GPT35TurboMaxTokens,
    "gemini-1.5-flash": Gemini15FlashMaxTokens,
    "gemini-1.5-pro":   Gemini15ProMaxTokens,
    "gemini-1.0-pro":   Gemini10ProMaxTokens,
}

func RunInputLimitsCheck(prompts []string, filenames []string, cfg *config.Config) (string, error) {
	for i, promptText := range prompts {
		model := cost.GetModel(promptText, cfg)
		nofTokens := cost.GetNumTokensFromPrompt(promptText, cfg)
		errOnLimits := checkIfTokensExceedsLimits(nofTokens, model)
		if errOnLimits != nil {
			problem := filenames[i] + "," + model
			return problem, fmt.Errorf("error on input tokens limits: %v", errOnLimits)
		}
	}
	return "", nil
}

func checkIfTokensExceedsLimits(nofTokens int, model string) error {
    maxTokens, exists := ModelMaxTokens[model]
    if !exists {
        return fmt.Errorf("model '%s' not found", model)
    }
    if nofTokens > maxTokens {
        return fmt.Errorf("number of tokens in prompt (%d) exceeds limits for model '%s' (max allowed: %d)", nofTokens, model, maxTokens)
    }
    return nil
}

func RunUserCheck(totalCost string, cfg *config.Config) error {
	if  cfg.Project.LLM.Provider == "GoogleAI" {
		fmt.Println("Unless you are using a free tier with Google AI, the total cost (USD - $) to run this review is at least:", totalCost)
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
