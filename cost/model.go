package cost

import (
	"log"
	"prismAId/config"

	openai "github.com/sashabaranov/go-openai"
	anthropic "github.com/anthropics/anthropic-sdk-go"
)

func GetModel(prompt string, cfg *config.Config) string {
	var modelFunc func(string, *config.Config) string
	switch cfg.Project.LLM.Provider {
	case "OpenAI":
		modelFunc = getOpenAIModel
	case "GoogleAI":
		modelFunc = getGoogleAIModel
	case "Cohere":
		modelFunc = getCohereModel
	case "Anthropic":
		modelFunc = getAnthropicModel
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

func getCohereModel(prompt string, cfg *config.Config) string {
	model := "command-r"
	switch cfg.Project.LLM.Model {
	case "": 
		// cost optimization, command-r is currently the cheapest and with the most input tokens allowed
	case "command": // leave the model selected by the user, but chek if supported
		model = cfg.Project.LLM.Model
	case "command-light":
		model = cfg.Project.LLM.Model
	case "command-r":
		model = cfg.Project.LLM.Model
	case "command-r-plus":
		model = cfg.Project.LLM.Model
	default:
		log.Println("Unsopported model: ", cfg.Project.LLM.Model)
		return ""
	}
	return model
}

func getAnthropicModel(prompt string, cfg *config.Config) string {
	model := anthropic.ModelClaude_3_Haiku_20240307
	switch cfg.Project.LLM.Model {
	case "": // cost optimization
		// all models have the same context window size, hence leave to haiku as the cheapest
	case "claude-3-5-sonnet":
		model = anthropic.ModelClaude_3_5_Sonnet_20240620
	case "claude-3-opus":
		model = anthropic.ModelClaude_3_Opus_20240229
	case "claude-3-sonnet":
		model = anthropic.ModelClaude_3_Sonnet_20240229
	case "claude-3-haiku":
		model = anthropic.ModelClaude_3_Haiku_20240307
	default:
		log.Println("Unsopported model: ", cfg.Project.LLM.Model)
		return ""
	}
	return model
}