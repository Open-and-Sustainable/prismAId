package convert

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"prismAId/config"
)

func Convert(config *config.Config) error {
	// Load files from the input directory
	inputDir := config.Project.Configuration.InputDirectory
	files, err := os.ReadDir(inputDir)
	if err != nil {
		log.Println("Error: ", err)
		return fmt.Errorf("error reading input directory: %v", err)
	}
	// formats
	formats := strings.Split(config.Project.Configuration.InputConversion, ",")
	// parse files
	for format := range formats {
		for _, file := range files {
			fullPath := filepath.Join(inputDir, file.Name())

			if filepath.Ext(file.Name()) == "."+formats[format] {
				txt_content, err := readText(fullPath, formats[format])
				if err == nil {
					fileNameWithoutExt := strings.TrimSuffix(file.Name(), "."+formats[format])
					txtPath := filepath.Join(inputDir, fileNameWithoutExt+".txt")
					
					err = writeText(txt_content, txtPath)
					if err != nil {
						log.Println("Error: ", err)
						return fmt.Errorf("error writing to file: %v", err)
					}
				}
			} else if filepath.Ext(file.Name()) == ".htm" { // this is to treat the special case of html files svaed with .htm extension
				txt_content, err := readText(fullPath, "html")
				if err == nil {
					fileNameWithoutExt := strings.TrimSuffix(file.Name(), ".htm")
					txtPath := filepath.Join(inputDir, fileNameWithoutExt+".txt")
					err = writeText(txt_content, txtPath)
					if err != nil {
						log.Println("Error: ", err)
						return fmt.Errorf("error writing to file: %v", err)
					}
				}
			}
		}
	}
	return nil
}

func readText(file string, format string) (string, error) {
	var modelFunc func(string) (string, error)
	switch format {
	case "pdf":
		modelFunc = readPdf
	case "docx":
		modelFunc = readDocx
	case "html":
		modelFunc = readHtml
	default:
		log.Println("Unsupported document type: ", format)
		return "", fmt.Errorf("unsupported document type: %s", format)
	}
	return modelFunc(file)
}

// writeText writes the given text to the specified file path
func writeText(text string, txtPath string) error {
	// Open the file for writing. If the file doesn't exist, it will be created.
	// The os.O_WRONLY flag opens the file for writing, and os.O_CREATE creates the file if it doesn't exist.
	// os.O_TRUNC truncates the file if it already exists.
	file, err := os.OpenFile(txtPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("error opening or creating file: %v", err)
	}
	defer file.Close() // Ensure that the file is properly closed after writing

	// Write the text to the file
	_, err = file.WriteString(text)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	fmt.Printf("Successfully wrote to %s\n", txtPath)
	return nil
}