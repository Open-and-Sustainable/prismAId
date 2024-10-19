package config

import (
	"os"

	"github.com/BurntSushi/toml"
)


// Config defines the top-level configuration structure, matching the TOML file layout.
type Config struct {
	Project ProjectConfig         `toml:"project"`
	Prompt  PromptConfig          `toml:"prompt"`
	Review  map[string]ReviewItem `toml:"review"`
}

// ProjectConfig holds details about the project, its metadata, and settings.
type ProjectConfig struct {
	Name          string               `toml:"name"`
	Author        string               `toml:"author"`
	Version       string               `toml:"version"`
	Configuration ProjectConfiguration `toml:"configuration"`
	LLM           LLMConfig            `toml:"llm"`
}

// ProjectConfiguration defines various settings related to project input and output.
type ProjectConfiguration struct {
	InputDirectory  string `toml:"input_directory"`
	InputConversion string `toml:"input_conversion"`
	ResultsFileName string `toml:"results_file_name"`
	OutputFormat    string `toml:"output_format"`
	LogLevel        string `toml:"log_level"`
	Ensemble		string	`toml:"ensemble"`
	CotJustification string  `toml:"cot_justification"`
	Duplication      string  `toml:"duplication"`
	Summary    string     `toml:"summary"`
}

// LLMConfig holds the configuration settings specific to the AI model being used.
type LLMConfig struct {
	Provider       string  `toml:"provider"`
	ApiKey         string  `toml:"api_key"`
	Model          string  `toml:"model"`
	Temperature    float64 `toml:"temperature"`
	TpmLimit       int64   `toml:"tpm_limit"`
	RpmLimit       int64   `toml:"rpm_limit"`
}

// PromptConfig specifies the configurations related to task prompting.
type PromptConfig struct {
	Persona        string `toml:"persona"`
	Task           string `toml:"task"`
	ExpectedResult string `toml:"expected_result"`
	Failsafe       string `toml:"failsafe"`
	Definitions    string `toml:"definitions"`
	Example        string `toml:"example"`
}

// ReviewItem defines key-value pairs for review configurations.
type ReviewItem struct {
	Key    string   `toml:"key"`
	Values []string `toml:"values"`
}

// LoadConfig reads a TOML configuration file and returns a Config instance.
// If the configuration file does not specify certain values, defaults are applied.
//
// Parameters:
//   - path: A string representing the file path to the configuration file.
//
// Returns:
//   - A pointer to a Config struct containing the loaded configuration.
//   - An error if any issues occur while reading or parsing the configuration file.
//
// Example:
//   > config, err := LoadConfig("path/to/config.toml")
//   > if err != nil {
//   >     log.Fatal("Failed to load config:", err)
//   > }
func LoadConfig(path string) (*Config, error) {
	var config Config
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		return nil, err
	}

	// API keys
	if config.Project.LLM.ApiKey == "" {
		if config.Project.LLM.Provider == "OpenAI" {
			config.Project.LLM.ApiKey = os.Getenv("OPENAI_API_KEY")
		} else if config.Project.LLM.Provider == "GoogleAI" {
			config.Project.LLM.ApiKey = os.Getenv("GOOGLE_AI_API_KEY")
		} else if config.Project.LLM.Provider == "Cohere" {
			config.Project.LLM.ApiKey = os.Getenv("CO_API_KEY")
		} else if config.Project.LLM.Provider == "Anthropic" {
			config.Project.LLM.ApiKey = os.Getenv("ANTHROPIC_API_KEY")
		}
	}

	// Default values

	if config.Project.Configuration.InputConversion == "" {
		config.Project.Configuration.InputConversion = "no"
	}

	if config.Project.Configuration.OutputFormat == "" {
		config.Project.Configuration.OutputFormat = "csv"
	}

	if config.Project.Configuration.LogLevel == "" {
		config.Project.Configuration.LogLevel = "low"
	}

	if config.Project.LLM.Temperature == 0 {
		config.Project.LLM.Temperature = 0
	}

	if config.Project.Configuration.CotJustification == "" {
		config.Project.Configuration.CotJustification = "no"
	}

	if config.Project.Configuration.Summary == "" {
		config.Project.Configuration.Summary = "no"
	}

	if config.Project.Configuration.Duplication == "" {
		config.Project.Configuration.Duplication = "no"
	}

	if config.Project.LLM.TpmLimit == 0 {
		config.Project.LLM.TpmLimit = 0 // This would mean no delay is applied
	}

	if config.Project.LLM.RpmLimit == 0 {
		config.Project.LLM.RpmLimit = 0 // This would mean no delay is applied
	}

	return &config, nil
}
