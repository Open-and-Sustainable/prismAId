package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"prismAId/config"
	"prismAId/cost"
	"prismAId/llm"
	"prismAId/prompt"
	"prismAId/results"
	"strings"
	"time"
)

func main() {
	// Define a flag for the project configuration file
	projectConfigPath := flag.String("project", "", "Path to the project configuration file")

	// Parse the flags
	flag.Parse()

	// Check if the project flag is provided
	if *projectConfigPath == "" {
		fmt.Println("Usage: ./prismAId_OS_CPU[.exe] --project <path-to-your-project-config.toml>")
		os.Exit(1)
	}

	// load project configuration
	config, err := config.LoadConfig(*projectConfigPath)
	if err != nil {
		fmt.Println("Error loading project configuration:", err) // here the logging function is not implemented yet
		return
	}
	// setup logging
	if config.Project.Configuration.LogLevel == "high" {
		setupLogging(File, *projectConfigPath)
	} else if config.Project.Configuration.LogLevel == "medium" {
		setupLogging(Stdout, *projectConfigPath)
	} else {
		setupLogging(Silent, *projectConfigPath) // default value
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

	// start writer for results.. the file will be project_name[.csv or .json] in the path where the toml is
	resultsFileName := config.Project.Configuration.ResultsFileName
	outputFilePath := resultsFileName + "." + output_format
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		log.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close() // Ensure the file is closed after all operations are done
	if output_format == "json" {
		for i, prompt := range prompts {
			fmt.Printf("Processing file #%d/%d: %s\n", i+1, len(prompts), filenames[i])
			log.Println("File: ", filenames[i], " Prompt: ", prompt)
			response, err := llm.QueryLLM(prompt, config)
			if err != nil {
				log.Println("Error querying LLM:", err)
				return
			}
			results.WriteJSONData(response, outputFile) // Write formatted JSON to file
			// sleep before next prompt if it's not the last one
			if i < len(prompts)-1 {
				waitWithStatus(getWaitTime(prompt, config))
			}
		}
	} else {
		keys := prompt.GetResultsKeys(config)
		writer := results.CreateCSVWriter(outputFile, keys) // Pass the file to CreateWriter
		defer writer.Flush()                                // Ensure data is flushed after all writes
		// send prompts to API and write results on file
		for i, prompt := range prompts {
			fmt.Printf("Processing file #%d/%d: %s\n", i+1, len(prompts), filenames[i])
			log.Println("File: ", filenames[i], " Prompt: ", prompt)
			response, err := llm.QueryLLM(prompt, config)
			if err != nil {
				log.Println("Error querying LLM:", err)
				return
			}
			results.WriteCSVData(response, filenames[i], writer, keys)
			// sleep before next prompt if it's not the last one
			if i < len(prompts)-1 {
				waitWithStatus(getWaitTime(prompt, config))
			}
		}
	}
}

type LogLevel int

const (
	Silent LogLevel = iota
	Stdout
	File
)

// Setup logging based on log level
func setupLogging(level LogLevel, filename string) {
	var logOutput io.Writer
	switch level {
	case Silent:
		logOutput = io.Discard // Discard all log output
	case Stdout:
		logOutput = os.Stdout // Log to standard output
	case File:
		logname := strings.TrimSuffix(filename, ".toml") + ".log"
		logFile, err := os.OpenFile(logname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		logOutput = logFile // Log to file
	default:
		logOutput = io.Discard // Default to discarding output
	}

	log.SetOutput(logOutput)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// Method that returns the number of seconds to wait to respect TPM limits
func getWaitTime(prompt string, config *config.Config) int {
	// Get the number of tokens from the prompt
	tokens := cost.GetNumTokensFromPrompt(prompt, config)
	// Get the TPM limit from the configuration
	tpmLimit := config.Project.LLM.TpmLimit
	// Calculate the time to wait until the next minute
	now := time.Now()
	remainingSeconds := 60 - now.Second()
	// Calculate the number of tokens per second allowed
	tokensPerSecond := float64(tpmLimit) / 60.0
	// Calculate the required wait time in seconds to not exceed TPM limit
	requiredWaitTime := float64(tokens) / tokensPerSecond
	// Calculate the seconds to the next minute
	secondsToMinute := 0
	if int(requiredWaitTime) > 60 {
		secondsToMinute = 60 - int(requiredWaitTime)%60
	}
	// If required wait time is more than remaining seconds in the current minute, wait until next minute
	if requiredWaitTime > float64(remainingSeconds) {
		return remainingSeconds + int(requiredWaitTime) + secondsToMinute
	}
	// Otherwise, calculate the wait time based on tokens used
	return remainingSeconds
}

func waitWithStatus(waitTime int) {
	ticker := time.NewTicker(1 * time.Second) // Ticks every second
	defer ticker.Stop()
	remainingTime := waitTime
	for range ticker.C {
		// Print the status only when the remaining time modulo 5 equals 0
		if remainingTime%5 == 0 {
			fmt.Printf("Waiting... %d seconds remaining\n", remainingTime)
		}
		remainingTime--
		// Break the loop when no time is left
		if remainingTime <= 0 {
			fmt.Println("Wait completed.")
			break
		}
	}
}
