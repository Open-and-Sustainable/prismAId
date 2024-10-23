package config

import (
	"os"

	"github.com/BurntSushi/toml"
)

// FileReader is an interface for reading files.
type FileReader interface {
    ReadFile(path string) ([]byte, error)
}

// EnvReader is an interface for accessing environment variables.
type EnvReader interface {
    GetEnv(key string) string
}

type RealFileReader struct{}

func (r RealFileReader) ReadFile(path string) ([]byte, error) {
    return os.ReadFile(path)
}

type RealEnvReader struct{}

func (r RealEnvReader) GetEnv(key string) string {
    return os.Getenv(key)
}

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
	LLM           map[string]LLMItem   `toml:"llm"`
}

// ProjectConfiguration defines various settings related to project input and output.
type ProjectConfiguration struct {
	InputDirectory  string `toml:"input_directory"`
	InputConversion string `toml:"input_conversion"`
	ResultsFileName string `toml:"results_file_name"`
	OutputFormat    string `toml:"output_format"`
	LogLevel        string `toml:"log_level"`
	CotJustification string  `toml:"cot_justification"`
	Duplication      string  `toml:"duplication"`
	Summary    string     `toml:"summary"`
}

// LLMConfig holds the configuration settings specific to the AI model being used.
type LLMItem struct {
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
//func LoadConfig(path string) (*Config, error) {
func LoadConfig(path string, fileReader FileReader, envReader EnvReader) (*Config, error) {
	var config Config

    // Read the file using the injected FileReader interface
    data, err := fileReader.ReadFile(path)
    if err != nil {
        return nil, err
    }

    // Decode the TOML data
    if _, err := toml.Decode(string(data), &config); err != nil {
        return nil, err
    }

	for key, llm := range config.Project.LLM {
		if llm.ApiKey == "" {  // If API key is empty, look for it in environment variables
			switch llm.Provider {
			case "OpenAI":
				llm.ApiKey = envReader.GetEnv("OPENAI_API_KEY")
			case "GoogleAI":
				llm.ApiKey = envReader.GetEnv("GOOGLE_AI_API_KEY")
			case "Cohere":
				llm.ApiKey = envReader.GetEnv("CO_API_KEY")
			case "Anthropic":
				llm.ApiKey = envReader.GetEnv("ANTHROPIC_API_KEY")
			}
		}

		if llm.Temperature < 0 {
			llm.Temperature = 0
		}
		if llm.TpmLimit < 0 {
			llm.TpmLimit = 0
		}
		if llm.RpmLimit < 0 {
			llm.RpmLimit = 0 
		}
		// Update the map directly with the modified llm
		config.Project.LLM[key] = llm
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

	if config.Project.Configuration.CotJustification == "" {
		config.Project.Configuration.CotJustification = "no"
	}

	if config.Project.Configuration.Summary == "" {
		config.Project.Configuration.Summary = "no"
	}

	if config.Project.Configuration.Duplication == "" {
		config.Project.Configuration.Duplication = "no"
	}

	return &config, nil
}
