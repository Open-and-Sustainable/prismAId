package check

import (
	"fmt"

	"prismAId/config"
	"prismAId/cost"

	anthropic "github.com/anthropics/anthropic-sdk-go"
	openai "github.com/sashabaranov/go-openai"
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
	// Cohere Models
	CommandMaxTokens = 4096
	CommandLightMaxTokens = 4096
	CommandRMaxTokens = 128000
	CommandRPlusMaxTokens = 128000
    // Anthropic Models
    AnthropicMaxTokens = 200000
)

var ModelMaxTokens = map[string]int{
    openai.GPT4oMini:   GPT4MiniMaxTokens,
    openai.GPT4o:       GPT4MaxTokens,
    openai.GPT4Turbo:   GPT4TurboMaxTokens,
    openai.GPT3Dot5Turbo:    GPT35TurboMaxTokens,
    "gemini-1.5-flash": Gemini15FlashMaxTokens,
    "gemini-1.5-pro":   Gemini15ProMaxTokens,
    "gemini-1.0-pro":   Gemini10ProMaxTokens,
	"command-r-plus":   CommandRPlusMaxTokens,
	"command-r":        CommandRMaxTokens,
	"command-light":    CommandLightMaxTokens,
	"command":          CommandMaxTokens,
    anthropic.ModelClaude_3_5_Sonnet_20240620:      AnthropicMaxTokens,
    anthropic.ModelClaude_3_Sonnet_20240229:      AnthropicMaxTokens,
    anthropic.ModelClaude_3_Opus_20240229:      AnthropicMaxTokens,
    anthropic.ModelClaude_3_Haiku_20240307:      AnthropicMaxTokens,
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

