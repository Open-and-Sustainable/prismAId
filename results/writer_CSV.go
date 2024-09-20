package results

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

//type ResponseData map[string]interface{}

// CreateWriter writes a header based on keys sorted alphabetically
func CreateCSVWriter(outputFile *os.File, keys []string) *csv.Writer {
	writer := csv.NewWriter(outputFile)
	// Prepare header
	header := make([]string, len(keys)+1)
	header[0] = "File Name"
	for i, key := range keys {
		header[i+1] = key
	}
	writer.Write(header)
	return writer
}

// WriteCSVData writes the parsed JSON data to CSV, handling both single objects and arrays of objects
func WriteCSVData(response string, fileNameWithoutExt string, writer *csv.Writer, keys []string) {
	// Clean the response
	response = cleanJSON(response)

	// Unmarshal the cleaned JSON string into an interface to detect the structure
	var rawData interface{}
	err := json.Unmarshal([]byte(response), &rawData)
	if err != nil {
		log.Println("Error parsing JSON:", err)
		return
	}

	// Check if rawData is a slice (array) or a map (single object)
	switch data := rawData.(type) {
	case []interface{}:
		// If it's an array, process each object in the array
		for _, item := range data {
			obj, ok := item.(map[string]interface{})
			if !ok {
				log.Println("Error: invalid object in array")
				continue
			}
			writeObjectToCSV(obj, fileNameWithoutExt, writer, keys)
		}
	case map[string]interface{}:
		// If it's a single object, process it directly
		writeObjectToCSV(data, fileNameWithoutExt, writer, keys)
	default:
		log.Println("Error: unexpected JSON structure")
	}
}

// writeObjectToCSV processes a single object and writes it to the CSV
func writeObjectToCSV(obj map[string]interface{}, fileNameWithoutExt string, writer *csv.Writer, keys []string) {
	// Prepare data row for CSV
	row := make([]string, len(keys)+1) // One extra field for filename
	row[0] = fileNameWithoutExt         // First column is the filename

	// Loop through the keys and write each corresponding value to the CSV
	for i, key := range keys {
		// First, try the exact key as specified
		value, exists := obj[key]
		if !exists {
			// If not found, try replacing spaces with underscores in the key
			keyWithUnderscores := strings.ReplaceAll(key, " ", "_")
			value, exists = obj[keyWithUnderscores]
		}

		// If the key (or its underscore version) does not exist, leave the column blank
		if !exists {
			row[i+1] = ""
			continue
		}

		// Handle different types of values (string, array, etc.)
		switch v := value.(type) {
		case string:
			row[i+1] = v // If it's a string, use it directly
		case []interface{}:
			// If it's an array, convert each element to a string and join them with ";"
			var strValues []string
			for _, elem := range v {
				strValues = append(strValues, fmt.Sprintf("%v", elem)) // Convert each element to string
			}
			row[i+1] = strings.Join(strValues, "; ") // Join array elements with "; "
		default:
			row[i+1] = fmt.Sprintf("%v", v) // Handle other types (numbers, bools, etc.)
		}
	}

	// Write the row to CSV
	if err := writer.Write(row); err != nil {
		log.Println("Error writing to CSV:", err)
		return
	}

	// Flush the writer and check for any errors
	writer.Flush()
	if err := writer.Error(); err != nil {
		log.Println("Error flushing CSV:", err)
	}
}