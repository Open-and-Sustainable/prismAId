package debug

import (
	"log"
	"os"
	"path/filepath"
	"prismAId/config"
	"strings"
)

const duplication_extension = "duplicate"

// DuplicateInput reads all text files from the configured input directory and creates copies of them with a 
// specified duplication extension. This function is useful for creating backup copies of input data or for 
// testing purposes.
//
// Arguments:
// - config: A pointer to the application’s configuration which holds the input directory details.
//
// Returns:
// - An error if the directory cannot be read or if a file operation fails, otherwise returns nil.
func DuplicateInput(config *config.Config) error {
	// Load text files from the input directory
	files, err := os.ReadDir(config.Project.Configuration.InputDirectory)
	if err != nil {
		log.Fatal(err)
		return err
	}

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

			// Create the new filename with the duplication extension
			fileBaseName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
			newFileName := fileBaseName + "_" + duplication_extension + ".txt"
			newFilePath := filepath.Join(config.Project.Configuration.InputDirectory, newFileName)

			// Write the duplicated content to the new file
			err = os.WriteFile(newFilePath, content, 0644)
			if err != nil {
				log.Printf("Failed to write duplicated file %s: %v", newFileName, err)
				return err
			}

			log.Printf("File %s duplicated as %s", file.Name(), newFileName)
		}
	}

	return nil
}

func RemoveDuplicateInput(config *config.Config) error {
	// Load files from the input directory
	files, err := os.ReadDir(config.Project.Configuration.InputDirectory)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Iterate over each file in the directory
	for _, file := range files {
		// Check if the file is a .txt file
		if filepath.Ext(file.Name()) == ".txt" {
			fileBaseName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
			expectedSuffix := "_" + duplication_extension

			// If the filename ends with the duplication extension and .txt, it's a duplicate
			if strings.HasSuffix(fileBaseName, expectedSuffix) {
				// Construct the full file path
				filePath := filepath.Join(config.Project.Configuration.InputDirectory, file.Name())

				// Remove the file
				err := os.Remove(filePath)
				if err != nil {
					log.Printf("Failed to remove file %s: %v", file.Name(), err)
					return err
				}

				log.Printf("Removed duplicated file: %s", file.Name())
			}
		}
	}

	return nil
}