package llm

import (
	"fmt"
	"prismAId/config"
)

func QueryLLM(prompt string, cfg *config.Config) (string, error) {
	var queryFunc func(string, *config.Config) (string, error)

	switch cfg.Project.LLM.Provider {
	case "OpenAI":
		queryFunc = queryOpenAI
	case "GoogleAI":
		queryFunc = queryGoogleAI
	default:
		return "", fmt.Errorf("unsupported LLM provider: %s", cfg.Project.LLM.Provider)
	}

	return queryFunc(prompt, cfg)
}
