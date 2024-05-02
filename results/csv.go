package results

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

type ResponseData struct {
	Period          string `json:"period (years)"`
	AnnualRateBasis string `json:"annual rate (basis points)"`
}

// CreateWriter accepts an already opened file and returns a configured csv.Writer
func CreateWriter(outputFile *os.File) *csv.Writer {
	writer := csv.NewWriter(outputFile)
	writer.Write([]string{"File Name", "Period (Year)", "Annual Rate (Basis Points)"})
	return writer
}

func WriteData(response string, fileNameWithoutExt string, writer *csv.Writer) {
	// Parse JSON response
	var data ResponseData
	err := json.Unmarshal([]byte(response), &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
	}
	//fmt.Println("Data period (years):", data.Period)
	//fmt.Println("Data annual rate (basis points):", data.AnnualRateBasis)

	// Write data to CSV
	if err := writer.Write([]string{fileNameWithoutExt, data.Period, data.AnnualRateBasis}); err != nil {
		fmt.Println("Error writing to CSV:", err)
		return // Handle write error
	}

	// Explicitly flush data to handle any potential buffering issues
	writer.Flush()
	if err := writer.Error(); err != nil {
		fmt.Println("Error flushing CSV:", err)
		return // Handle flush error
	}
}
