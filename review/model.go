package review

import (
	"github.com/open-and-sustainable/prismaid/config"
)

type Model struct {
	Provider 	string
	Model    	string
	APIKey   	string
	Temperature	float64
	TPM			int64
	RPM			int64
	ID			string
}

// NewModels constructs a slice of Model structures from a given map of LLMItem pointers.
//
// Parameters:
//   llmMap - A map where each key is a string and each value is a pointer to an LLMItem,
//            representing different language learning model configurations.
//
// Returns:
//   []Model - A slice of Model structures. Each Model contains the provider, model name,
//             API key, temperature, tokens per minute limit, requests per minute limit,
//             and the key from the map as its identifier.
//
//   error - Returns nil as no errors are expected under normal operation of this function.
//           If future implementations introduce error scenarios, this return value can
//           be used to handle them accordingly.
//
// The function iterates over the llmMap, and for each entry, it constructs a Model struct
// and appends it to the slice of Models. The key from the map is used as an identifier in
// the resulting Model structure. The function currently always returns nil for the error
// since no error conditions are handled within it. This setup allows for easy addition
// of error handling in future without changing the function signature.
func NewModels(llmMap map[string]config.LLMItem) ([]Model, error) {
	var models []Model
	for key, value := range llmMap {
		models = append(models, Model{value.Provider, value.Model, value.ApiKey, value.Temperature, value.TpmLimit, value.RpmLimit, key}) 
	}
	return models, nil
}