package cost

import (
	"log"
	"prismAId/config"

	"github.com/sashabaranov/go-openai"
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

func GetModel(prompt string, cfg *config.Config) string {
	var modelFunc func(string, *config.Config) string
	switch cfg.Project.LLM.Provider {
	case "OpenAI":
		modelFunc = getOpenAIModel
	case "GoogleAI":
		modelFunc = getGoogleAIModel
	default:
		log.Println("Unsupported LLM provider: ", cfg.Project.LLM.Provider)
		return ""
	}
	return modelFunc(prompt, cfg)
}

func getOpenAIModel(prompt string, cfg *config.Config) string {
	model := openai.GPT4oMini
	switch cfg.Project.LLM.Model {
	case "": // cost optimization
		// old code before GPT 4 Omni mini model availability -- now the only solution to minimize the cost
		/*numTokens := numTokensFromMessages([]openai.ChatCompletionMessage{{Role: openai.ChatMessageRoleUser, Content: prompt}}, model)
		if numTokens > 16385 {
			model = openai.GPT4o
		}*/
	case "gpt-3.5-turbo":
		model = openai.GPT3Dot5Turbo
	case "gpt-4-turbo":
		model = openai.GPT4Turbo
	case "gpt-4o":
		model = openai.GPT4o
	case "gpt-4o-mini":
		model = openai.GPT4oMini
	default:
		log.Println("Unsopported model: ", cfg.Project.LLM.Model)
		return ""
	}
	return model
}

func getGoogleAIModel(prompt string, cfg *config.Config) string {
	model := "gemini-1.0-pro"
	switch cfg.Project.LLM.Model {
	case "": // cost optimization, input token limit values: gemini-1.0-pro 30720, gemini-1.5-flash 1048576, gemini-1.5-pro 2097152
		numTokens := numTokensFromPromptGoogleAI(prompt, model, cfg)
		if numTokens > 30720 && numTokens <= 1048576 {
			model = "gemini-1.5-flash"
		} else if numTokens > 1048576 {
			model = "gemini-1.5-pro"
		}
	case "gemini-1.0-pro": // leave the model selected by the user, but chek if supported
		model = cfg.Project.LLM.Model
	case "gemini-1.5-flash":
		model = cfg.Project.LLM.Model
	case "gemini-1.5-pro":
		model = cfg.Project.LLM.Model
	default:
		log.Println("Unsopported model: ", cfg.Project.LLM.Model)
		return ""
	}
	return model
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
	default:
		log.Println("Unsupported LLM provider: ", cfg.Project.LLM.Provider)
		return 0
	}
	return numTokens
}
