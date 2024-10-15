package llm

import (
	"fmt"
	"prismAId/config"
)

const justification_query = "For each one of the keys and answers you provided, provide a justification for your answer as a chain of thought. In particular, I want a textual description of the few stages of the chin of thought that lead you to the answer you provided and the sentences in the text you analyzes that support your decision. If the value of a key was 'no' or empty '' because of lack of information on that topic in the text analyzed, explicitly report this reason. Please provide only th einformation requested, neither introductory nor concluding remarks."
const summary_query = "Summarize in very few sentences the text provided before for your review."

func QueryLLM(prompt string, cfg *config.Config) (string, string, string, error) {
	var queryFunc func(string, *config.Config) (string, string, string, error)

	switch cfg.Project.LLM.Provider {
	case "OpenAI":
		queryFunc = queryOpenAI
	case "GoogleAI":
		queryFunc = queryGoogleAI
	case "Cohere":
		queryFunc = queryCohere
	default:
		return "", "", "", fmt.Errorf("unsupported LLM provider: %s", cfg.Project.LLM.Provider)
	}

	return queryFunc(prompt, cfg)
}
