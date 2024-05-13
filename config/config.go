package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

// Define your configuration structures matching the TOML structure
type Config struct {
	Project ProjectConfig         `toml:"project"`
	Prompt  PromptConfig          `toml:"prompt"`
	Review  map[string]ReviewItem `toml:"review"`
}

type ProjectConfig struct {
	Name          string               `toml:"name"`
	Author        string               `toml:"author"`
	Version       string               `toml:"version"`
	Configuration ProjectConfiguration `toml:"configuration"`
	LLM           LLMConfig            `toml:"llm"`
}

type ProjectConfiguration struct {
	InputDirectory   string `toml:"input_directory"`
	ResultsDirectory string `toml:"results_directory"`
	OutputFormat     string `toml:"output_format"`
	LogLevel         string `toml:"log_level"`
}

type LLMConfig struct {
	Provider       string  `toml:"provider"`
	ApiKey         string  `toml:"api_key"`
	Model          string  `toml:"model"`
	Temperature    float64 `toml:"temperature"`
	BatchExecution string  `toml:"batch_execution"`
}

type PromptConfig struct {
	Persona        string `toml:"persona"`
	Task           string `toml:"task"`
	ExpectedResult string `toml:"expected_result"`
	Failsafe       string `toml:"failsafe"`
	Example        string `toml:"example"`
}

type ReviewItem struct {
	Key    string   `toml:"key"`
	Values []string `toml:"values"`
}

func LoadConfig(path string) (*Config, error) {
	var config Config
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		return nil, err
	}
	if config.Project.LLM.ApiKey == "" {
		config.Project.LLM.ApiKey = os.Getenv("OPENAI_API_KEY")
	}
	return &config, nil
}
