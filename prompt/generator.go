// Prompt generation
package prompt

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"prismAId/config"
	"sort"
	"strings"
)

// ParsePrompts reads the configuration and generates a list of prompts along with their corresponding filenames.
// The function combines different parts of the prompts to create a structured list of inputs.
//
// Arguments:
// - config: A pointer to the application's configuration which specifies how prompts should be parsed and organized.
//
// Returns:
// - Two slices of strings: 
//   - The first slice contains the generated prompts.
//   - The second slice contains the filenames associated with each prompt.
func ParsePrompts(config *config.Config) ([]string, []string) {
	// This slice will store all combined prompts
	var prompts []string
	// This slice will store the filenames corresponding to each prompt
	var filenames []string

	// The common part of prompts
	expected_result := parseExpectedResults(config)
	common_part := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s",
		config.Prompt.Persona, config.Prompt.Task, expected_result,
		config.Prompt.Failsafe, config.Prompt.Definitions, config.Prompt.Example)

	// Load text files
	files, err := os.ReadDir(config.Project.Configuration.InputDirectory)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".txt" {
			filePath := filepath.Join(config.Project.Configuration.InputDirectory, file.Name())
			documentText, err := os.ReadFile(filePath)
			if err != nil {
				log.Println("Error reading file:", err)
				return nil, nil
			}

			// Combine prompt elements
			prompt := fmt.Sprintf("%s \n\n%s", common_part, documentText)
			// Append the combined text to the slice
			prompts = append(prompts, prompt)

			// Get the filename without extension
			fileNameWithoutExt := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
			// Append the filename to the filenames slice
			filenames = append(filenames, fileNameWithoutExt)
		}
	}

	return prompts, filenames
}

func parseExpectedResults(config *config.Config) string {
	expectedResult := config.Prompt.ExpectedResult
	keys := GetResultsKeysOrdered(config)

	// Build a map from sorted keys using descriptive keys
	sortedReviewItems := make(map[string][]string)
	for _, numericKey := range keys {
		item := config.Review[numericKey]
		sortedReviewItems[item.Key] = item.Values // Use the descriptive key for the JSON output
	}

	// Convert sorted map to JSON
	reviewJSON, err := json.Marshal(sortedReviewItems)
	if err != nil {
		log.Fatalf("Error marshalling review items to JSON: %v", err)
	}

	// Combine the expected result with the JSON-formatted review items
	fullSummary := fmt.Sprintf("%s %s", expectedResult, string(reviewJSON))
	return fullSummary
}

// GetResultsKeysOrdered retrieves the keys from the results configuration in a specific, sorted order.
// This function ensures that the keys are returned in a consistent order, which is useful for generating 
// organized outputs.
//
// Arguments:
// - config: A pointer to the application's configuration that specifies the result keys to be retrieved.
//
// Returns:
// - A slice of strings containing the ordered result keys.
func GetResultsKeysOrdered(config *config.Config) []string {
	// Collect keys for sorting based on numeric keys to maintain order
	keys := make([]string, 0, len(config.Review))
	for key := range config.Review {
		keys = append(keys, key)
	}
	sort.Strings(keys) // Sort keys to ensure order
	return keys
}

// GetResultsKeys retrieves the keys from the results configuration without enforcing a specific order.
// This function is useful when the order of the keys is not critical.
//
// Arguments:
// - config: A pointer to the application's configuration that specifies the result keys to be retrieved.
//
// Returns:
// - A slice of strings containing the result keys.
func GetResultsKeys(config *config.Config) []string {
	var keys []string
	for _, item := range config.Review {
		keys = append(keys, item.Key)
	}
	sort.Strings(keys) // Optional: Sort keys alphabetically for consistent output
	return keys
}
