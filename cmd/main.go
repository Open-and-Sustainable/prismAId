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
	// start writer for results
	out_directory := "../projects/output/test"
	outputFilePath := filepath.Join(out_directory, "output.csv")
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		log.Println("Error creating CSV file:", err)
		return
	}
	defer outputFile.Close() // Ensure the file is closed after all operations are done
	keys := prompt.GetResultsKeys(config)
	writer := results.CreateWriter(outputFile, keys) // Pass the file to CreateWriter
	defer writer.Flush()                             // Ensure data is flushed after all writes
	// send prompts to API and write results on file
	for i, prompt := range prompts {
		response, err := llm.QueryOpenAI(prompt, config)
		if err != nil {
			log.Println("Error querying OpenAI:", err)
			return
		}
		results.WriteData(response, filenames[i], writer, keys)
	}
}
