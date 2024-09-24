package init

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	prompt "github.com/cqroot/prompt"
	choose "github.com/cqroot/prompt/choose"
	input "github.com/cqroot/prompt/input"
)

// Function to interactively collect project configuration information with advanced prompt features
func RunInteractiveConfigCreation() {
	fmt.Println("Running interactive project configuration initialization...")

	// Ask for file path to save the configuration
	filePath, err := prompt.New().Ask("Enter file path to save the configuration:").Input(
		"./config.toml", input.WithHelp(true), input.WithValidateFunc(validatePath))
	CheckErr(err)
	
	// Prompt for project name with help text
	projectName, err := prompt.New().Ask("Enter project name:").Input(
		"Test project",
		input.WithHelp(true),
	)
	CheckErr(err)

	// Prompt for author name with help
	author, err := prompt.New().Ask("Enter author name:").Input(
		"Name Lastname",
		input.WithHelp(true),
	)
	CheckErr(err)

	// Prompt for version
	version, err := prompt.New().Ask("Enter project version:").Input(
		"0.1",
		input.WithHelp(true),
	)
	CheckErr(err)

	// Configuration details with help for each choice
	inputDir, err := prompt.New().Ask("Enter input directory (must exist):").Input(
		"./", 
		input.WithHelp(true), input.WithValidateFunc(validateDirectory))
	CheckErr(err)

	resultsFileName, err := prompt.New().Ask("Enter results directory (must exist):").Input(
		"./", 
		input.WithHelp(true), input.WithValidateFunc(validateDirectory))
	CheckErr(err)

	// Advanced choose with help
	outputFormat, err := prompt.New().Ask("Choose output format:").
		AdvancedChoose(
			[]choose.Choice{
				{Text: "csv", Note: "Comma-separated values format for easier readability."},
				{Text: "json", Note: "JavaScript Object Notation format for structured data."},
			},
			choose.WithHelp(true),)
	CheckErr(err)

	// Log level with help 
	logLevel, err := prompt.New().Ask("Choose log level:").
		AdvancedChoose(
			[]choose.Choice{
				{Text: "low", Note: "Low verbosity: minimal logging."},
				{Text: "medium", Note: "High verbosity: logs displayed on stdout."},
				{Text: "high", Note: "High verbosity: logs saved to a file for detailed review."},
			},
			choose.WithHelp(true),)
	CheckErr(err)

	// Duplication option with help
	duplication, err := prompt.New().Ask("Enable duplication (for debugging)?").
		AdvancedChoose(
			[]choose.Choice{
				{Text: "no", Note: "Do not duplicate reviews."},
				{Text: "yes", Note: "Duplicate the manuscripts to review, and the cost, useful for consistency checks."},
			},
			choose.WithHelp(true),)
	CheckErr(err)

	// Chain-of-thought justification option
	cotJustification, err := prompt.New().Ask("Enable chain-of-thought justification (saved on file)?").
	AdvancedChoose(
		[]choose.Choice{
			{Text: "no", Note: "Do not enable chain-of-thought justification."},
			{Text: "yes", Note: "Enable model justification for the answers in terms of chain of thought."},
		},
		choose.WithHelp(true),)
	CheckErr(err)

	// LLM provider selection with help
	provider, err := prompt.New().Ask("Choose LLM provider:").
		AdvancedChoose(
			[]choose.Choice{
				{Text: "OpenAI", Note: "OpenAI's GPT-3 or GPT-4 models."},
				{Text: "GoogleAI", Note: "GoogleAI's Gemini models."},
				{Text: "Cohere", Note: "Cohere's language models."},
			},
			choose.WithHelp(true),)
	CheckErr(err)

	// Prompt for API key with input mask (for security)
	apiKey, err := prompt.New().Ask("Enter LLM API key (leave it empty to use environment variable):").Input("", input.WithEchoMode(input.EchoPassword))
	CheckErr(err)

	// Model choice for the selected LLM provider
	model := ""
	if provider == "OpenAI" {
		model, err = prompt.New().Ask("Enter model to be used:").AdvancedChoose(
			[]choose.Choice{
				{Text: "", Note: "Model chosen automatically to minimize costs."},
				{Text: "gpt-3.5-turbo", Note: "GPT-3.5 Turbo."},
				{Text: "gpt-4-turbo", Note: "GPT-4 Turbo."},
				{Text: "gpt-4o", Note: "GPT-4 Omni."},
				{Text: "gpt-4o-mini", Note: "GPT-4 Omni Mini."},
			},
			choose.WithHelp(true),)

	} else if provider == "GoogleAI" {
		model, err = prompt.New().Ask("Enter model to be used:").AdvancedChoose(
			[]choose.Choice{
				{Text: "", Note: "Model chosen automatically to minimize costs."},
				{Text: "gemini-1.0-pro", Note: "Gemini 1.0 Pro."},
				{Text: "gemini-1.5-pro", Note: "Gemini 1.5 Pro."},
				{Text: "gemini-1.5-flash", Note: "Gemini 1.5 Flash."},
			},
			choose.WithHelp(true),)
	} else if provider == "Cohere" {
		model, err = prompt.New().Ask("Enter model to be used:").AdvancedChoose(
			[]choose.Choice{
				{Text: "", Note: "Model chosen automatically to minimize costs."},
				{Text: "command", Note: "Command."},
				{Text: "command-light", Note: "Command Light."},
				{Text: "command-r", Note: "Command R."},
				{Text: "command-r-plus", Note: "Command R+."},
			},
			choose.WithHelp(true),)
	}
	CheckErr(err)

	// Generate TOML config from user inputs
	config := generateTomlConfig(
		projectName, author, version,
		inputDir, resultsFileName, outputFormat, logLevel,
		duplication, cotJustification, provider, apiKey, model,
	)

	// Write the configuration to file
	err = writeTomlConfigToFile(config, filePath)
	if err != nil {
		fmt.Println("Error writing configuration file:", err)
	} else {
		fmt.Println("Configuration file created successfully at:", filePath)
	}
}

// Helper function to generate the TOML configuration string
func generateTomlConfig(projectName, author, version, inputDir, resultsFileName, outputFormat, logLevel, duplication, cotJustification, provider, apiKey, model string) string {
	config := fmt.Sprintf(`
[project]
name = "%s"
author = "%s"
version = "%s"

[project.configuration]
input_directory = "%s"
results_file_name = "%s"
output_format = "%s"
log_level = "%s"
duplication = "%s"
cot_justification = "%s"

[project.llm]
provider = "%s"
api_key = "%s"
model = "%s"
`, projectName, author, version, inputDir, resultsFileName, outputFormat, logLevel, duplication, cotJustification, provider, apiKey, model)
	return strings.TrimSpace(config)
}

// Helper function to write the TOML configuration to a file
func writeTomlConfigToFile(config, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(config)
	return err
}

// CheckErr is a helper to check and handle errors
func CheckErr(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

// validatePath checks if the directory in the given path exists and if the file name is valid.
func validatePath(path string) error {
	// Separate the directory from the file
	dir := filepath.Dir(path)
	file := filepath.Base(path)

	// Check if the directory path is valid
	if err := validateDirectory(dir); err != nil {
		return err
	}

	// Check if the file name contains invalid characters
	if err := validateFileName(file); err != nil {
		return err
	}

	// Path is valid
	return nil
}

//	validateDirectory checks if the given directory is valid.
func validateDirectory(dir string) error {
	// Check if the directory exists and is a valid directory
	info, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return fmt.Errorf("%s: %w", dir, fmt.Errorf("invalid path"))
	} else if err != nil {
		return err
	}

	if !info.IsDir() {
		return fmt.Errorf("%s is not a directory", dir)
	}

	// Directory is valid
	return nil
}

// validateFileName checks if the file name contains invalid characters
func validateFileName(fileName string) error {
	// Define a regular expression for invalid characters in file names
	// For example, on Windows: <>:"/\|?*
	invalidChars := regexp.MustCompile(`[<>:"/\\|?*]`)

	if invalidChars.MatchString(fileName) {
		return fmt.Errorf("%s: %w", fileName, fmt.Errorf("invalid filename"))
	}

	// You can also check for empty filenames or other restrictions, like file extension:
	if fileName == "" {
		return fmt.Errorf("filename cannot be empty: %w", fmt.Errorf("invalid filename"))
	}

	// Check .tom;:
	if !strings.HasSuffix(fileName, ".toml") {
	     return fmt.Errorf("filename must have a .toml extension")
	}

	return nil
}
