package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"prismAId/config"
	"prismAId/cost"
	"prismAId/llm"
	"prismAId/prompt"
	"prismAId/results"
)

func main() {
	// load project configuration
	config, err := config.LoadConfig("../projects/test.toml")
	if err != nil {
		fmt.Println("Error loading project configuration:", err) // here the logging function is not implemented yet
		return
	}
	// setup logging
	if config.Project.Configuration.LogLevel == "high" {
		setupLogging(filepath.ListSeparator)
	} else if config.Project.Configuration.LogLevel == "medium" {
		setupLogging(Stdout)
	} else {
		setupLogging(Silent) // default value
	}
	// generate prompts
	prompts, filenames := prompt.ParsePrompts(config)
	log.Println("Found", len(prompts), "files")
	// ask if continuing given the total cost
	check := cost.RunUserCheck(cost.ComputeCosts(prompts, config))
	if check != nil {
		log.Printf("Error:\n%v", check)
		os.Exit(0) // if the user stops the execution it is still a success run, hence exit code = 0, but the reason for the exit may be different hence is logged
	}
	// proceed with the execution
	// get desired output format
	output_format := "csv"
	if config.Project.Configuration.OutputFormat == "json" {
		output_format = "json"
	}
	// start writer for results
	out_directory := "../projects/output/test"
	outputFileName := "output" + "." + output_format
	outputFilePath := filepath.Join(out_directory, outputFileName)
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		log.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close() // Ensure the file is closed after all operations are done
	if output_format == "json" {
		for i, prompt := range prompts {
			log.Println("File: ", filenames[i], " Prompt: ", prompt)
			response, err := llm.QueryOpenAI(prompt, config)
			if err != nil {
				log.Println("Error querying OpenAI:", err)
				return
			}
			results.WriteJSONData(response, outputFile) // Write formatted JSON to file
		}
	} else {
		keys := prompt.GetResultsKeys(config)
		writer := results.CreateCSVWriter(outputFile, keys) // Pass the file to CreateWriter
		defer writer.Flush()                                // Ensure data is flushed after all writes
		// send prompts to API and write results on file
		for i, prompt := range prompts {
			log.Println("File: ", filenames[i], " Prompt: ", prompt)
			response, err := llm.QueryOpenAI(prompt, config)
			if err != nil {
				log.Println("Error querying OpenAI:", err)
				return
			}
			results.WriteCSVData(response, filenames[i], writer, keys)
		}
	}
}
