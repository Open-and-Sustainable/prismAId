package debug

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"prismAId/config"
	"prismAId/llm"
)

func Summarize(config *config.Config) error {
	// Load files from the input directory
	files, err := os.ReadDir(config.Project.Configuration.InputDirectory)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// create summary file writer
	summaryFilePath := filepath.Join(config.Project.Configuration.ResultsFileName, "_summary.txt")
	summaryFile, err := os.Create(summaryFilePath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer summaryFile.Close()

	// Iterate over each file in the directory
	for _, file := range files {
		// Process only .txt files
		if filepath.Ext(file.Name()) == ".txt" {
			// Construct the full file path
			filePath := filepath.Join(config.Project.Configuration.InputDirectory, file.Name())

			// Read the file content	
			content, err := os.ReadFile(filePath)
			if err != nil {
				log.Printf("Failed to read file %s: %v", file.Name(), err)
				return err
			}

			// Create the summary prompt
			summaryPrompt := fmt.Sprintf("Summarize the following text in exactly %s sentences:\n\n%s", string(config.Project.Configuration.SummaryLength), string(content))

			// Summarize the content
			summary, _, err := llm.QueryLLM(summaryPrompt, config)
			if err != nil {
				log.Printf("Failed to summarize file %s: %v", file.Name(), err)
				return err
			}

			// Write the file name and the summarized content to the summary file
			summaryFile.WriteString(file.Name() + "\n")
			summaryFile.WriteString(summary + "\n\n")
		}
	}

	return nil
}