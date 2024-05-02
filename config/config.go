package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Project ProjectConfig `toml:"project"`
	Inputs  InputConfig   `toml:"inputs"`
	LLM     LLMConfig     `toml:"llm"`
	Output  OutputConfig  `toml:"output"`
}

type ProjectConfig struct {
	Name    string `toml:"name"`
	Author  string `toml:"author"`
	Version string `toml:"version"`
}

type InputConfig struct {
	PDFDirectory   string `toml:"pdf_directory"`
	PromptTemplate string `toml:"prompt_template"`
}

type LLMConfig struct {
	Provider   string `toml:"provider"`
	APIKey     string `toml:"api_key"`
	ConfigFile string `toml:"config_file"`
	Endpoint   string `toml:"endpoint"`
}

type OutputConfig struct {
	ResultsDirectory string `toml:"results_directory"`
	OutputFormat     string `toml:"output_format"`
}

func LoadConfig(path string) (*Config, error) {
	var config Config
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
