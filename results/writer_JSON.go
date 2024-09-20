package results

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

// to start results as an array
func StartJSONArray(outputFile *os.File) error {
	_, err := outputFile.WriteString("[\n")
	if err != nil {
		log.Println("Error starting JSON array:", err)
		return err
	}
	return nil
}

// WriteJSONData writes the JSON data to the file.
func WriteJSONData(response string, filename string, outputFile *os.File) {
	// Strip out markdown code fences (```json ... ```) if present
	response = cleanJSON(response)

	// Unmarshal the JSON string into a map to modify it
	var data map[string]interface{}
	err := json.Unmarshal([]byte(response), &data)
	if err != nil {
		log.Println("Error unmarshaling JSON:", err)
		return
	}

	// Add the filename to the JSON data
	data["filename"] = filename

	// Marshal the modified data back into a JSON string
	modifiedJSON, err := json.MarshalIndent(data, "", "    ") // Indent for pretty JSON output
	if err != nil {
		log.Println("Error marshaling modified JSON:", err)
		return
	}

	// Write the modified JSON to the file
	_, err = outputFile.WriteString(string(modifiedJSON))
	if err != nil {
		log.Println("Error writing JSON to file:", err)
	}
}

func WriteCommaInJSONArray(outputFile *os.File) error {
	_, err := outputFile.WriteString(",\n")
	if err != nil {
		log.Println("Error writing comma in JSON array:", err)
		return err
	}
	return nil
}

func CloseJSONArray(outputFile *os.File) error {
	// Write the closing bracket
	_, err := outputFile.WriteString("\n]")
	if err != nil {
		log.Println("Error closing JSON array:", err)
		return err
	}

	return nil
}

// cleanJSON strips out the markdown code fences from the response if present.
func cleanJSON(response string) string {
	// Remove triple backticks and the "json" part (if present)
	response = strings.TrimPrefix(response, "```json")
	response = strings.TrimSuffix(response, "```")
	return strings.TrimSpace(response) // Trim any extra whitespace
}