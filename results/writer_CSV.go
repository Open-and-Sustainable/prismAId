package results

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"strings"
)

type ResponseData map[string]interface{}

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

// WriteCSVData writes the CSV data
func WriteCSVData(response string, fileNameWithoutExt string, writer *csv.Writer, keys []string) {
	var data ResponseData
	err := json.Unmarshal([]byte(response), &data)
	if err != nil {
		log.Println("Error parsing JSON:", err)
		return
	}

	// Prepare data row
	row := make([]string, len(keys)+1)
	row[0] = fileNameWithoutExt
	for i, key := range keys {
		switch value := data[key].(type) {
		case string:
			row[i+1] = value // Directly use the string
		case []interface{}:
			// Convert all interface{} elements to strings and concatenate them
			var strValues []string
			for _, v := range value {
				if str, ok := v.(string); ok {
					strValues = append(strValues, str)
				}
			}
			row[i+1] = strings.Join(strValues, "; ") // Join all strings with a semicolon and a space as delimiter
		default:
			row[i+1] = "" // Handle other types or missing key
		}
	}

	// Write data to CSV
	if err := writer.Write(row); err != nil {
		log.Println("Error writing to CSV:", err)
		return
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		log.Println("Error flushing CSV:", err)
	}
}
