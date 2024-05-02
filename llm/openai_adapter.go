// OpenAI specific implementation of the LLM interface
package llm

type OpenAIClient struct {
	apiKey   string
	endpoint string
}

func (oa *OpenAIClient) SendPrompt(prompt string) (string, error) {
	// Implementation for sending prompt to OpenAI and getting the response
}
