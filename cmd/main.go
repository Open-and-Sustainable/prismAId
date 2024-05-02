package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"prismAId/llm"
	"prismAId/results"
	"strings"
)

func main() {
	promptText := "Given the text at the bottom, identify and calculate any mentioned 'greenium' values. The 'greenium' represents the discount rate making green bonds cheaper to fund than similar conventional bonds. Extract data and present it in JSON format with keys for 'period (years)', which is the period to which the observation refers, and 'annual rate (basis points)'. If no specific numbers are mentioned, leave the values empty."
	exampleText := "For example, if a document states that the greenium is around -0.05 percentage points over the 2014-2019, the output should be: {\"period (years)\": \"2014-2019\", \"annual rate (basis points)\": \"-5\"}."
	//documentText := "In terms of the greenium, we have no evidence of a significant greenium. However, the point estimates of the greenium has an overall average around -7.07 bps. We observe that over the recent years, the greenium turned from slightly positive to negative. Particularly, the greenium dropped sharply in February 2020 and increased gradually from May 2020. This could be related to the COVID-19 stress on financial market during that time."

	// The directory for input output files
	directory := "../data"
	files, err := os.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	// the CSV file to outpout results
	outputFilePath := filepath.Join(directory, "output.csv")
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer outputFile.Close() // Ensure the file is closed after all operations are done

	writer := results.CreateWriter(outputFile) // Pass the file to CreateWriter
	defer writer.Flush()                       // Ensure data is flushed after all writes

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".txt" {
			filePath := filepath.Join(directory, file.Name())
			documentText, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}
			prompt := fmt.Sprintf("%s\n%s \n\n%s", promptText, exampleText, documentText) // Combine prompt, example, and text
			response, err := llm.QueryOpenAI(prompt)
			if err != nil {
				fmt.Println("Error querying OpenAI:", err)
				return
			}
			//fmt.Println("OpenAI response:", response)

			fileNameWithoutExt := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))

			results.WriteData(response, fileNameWithoutExt, writer)
		}
	}
}
