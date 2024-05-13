package results

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type ResponseData map[string][]string

// CreateWriter writes a header based on keys sorted alphabetically
func CreateWriter(outputFile *os.File, keys []string) *csv.Writer {
	writer := csv.NewWriter(outputFile)
	// Prepare header
	header := make([]string, len(keys)+1)
	header[0] = "File Name"
	for i, key := range keys {
		header[i+1] = fmt.Sprintf("%d (%s)", i+1, key)
	}
	writer.Write(header)
	return writer
}

// WriteData writes the CSV data
func WriteData(response string, fileNameWithoutExt string, writer *csv.Writer, keys []string) {
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
		if values, ok := data[key]; ok && len(values) > 0 {
			row[i+1] = values[0] // Use the first value, adjust as needed
		} else {
			row[i+1] = "" // No value for this key
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
