package review

import (
    "github.com/Open-and-Sustainable/prismaid/config"
    "reflect"
	"sort"
    "testing"
)

func TestNewModels(t *testing.T) {
    tests := []struct {
        name  string
        llmMap map[string]config.LLMItem
        want  []Model
        wantErr bool
    }{
        {
            name: "Single Item",
            llmMap: map[string]config.LLMItem{
                "model1": {Provider: "OpenAI", Model: "GPT-3", ApiKey: "somekey", Temperature: 0.2, TpmLimit: 100, RpmLimit: 50},
            },
            want: []Model{
                {Provider: "OpenAI", Model: "GPT-3", APIKey: "somekey", Temperature: 0.2, TPM: 100, RPM: 50, ID: "model1"},
            },
            wantErr: false,
        },
        {
            name: "Multiple Items",
            llmMap: map[string]config.LLMItem{
                "model1": {Provider: "OpenAI", Model: "GPT-3", ApiKey: "key1", Temperature: 0.2, TpmLimit: 100, RpmLimit: 50},
                "model2": {Provider: "GoogleAI", Model: "Bert", ApiKey: "key2", Temperature: 0.1, TpmLimit: 150, RpmLimit: 60},
            },
            want: []Model{
                {Provider: "OpenAI", Model: "GPT-3", APIKey: "key1", Temperature: 0.2, TPM: 100, RPM: 50, ID: "model1"},
                {Provider: "GoogleAI", Model: "Bert", APIKey: "key2", Temperature: 0.1, TPM: 150, RPM: 60, ID: "model2"},
            },
            wantErr: false,
        },
    }

    for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewModels(tt.llmMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewModels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("NewModels() got %d models, want %d models", len(got), len(tt.want))
				return
			}
			sortModels(got)
			sortModels(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewModels() got = %v, want %v", got, tt.want)
			}
		})
	}
	
}

// Define a sorting function for Model slices
func sortModels(models []Model) {
    sort.Slice(models, func(i, j int) bool {
        return models[i].ID < models[j].ID
    })
}
