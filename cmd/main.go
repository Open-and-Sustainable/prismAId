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
	"sync"
	"time"
)

// Global variable to store the timestamps of requests
var requestTimestamps []time.Time
var mutex sync.Mutex

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
	check := cost.RunUserCheck(cost.ComputeCosts(prompts, config), config)
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

	// Loop through the prompts
	for i, promptText := range prompts {
		fmt.Printf("Processing file #%d/%d: %s\n", i+1, len(prompts), filenames[i])
		log.Println("File: ", filenames[i], " Prompt: ", promptText)

		// Query the LLM
		response, justification, err := llm.QueryLLM(promptText, config)
		if err != nil {
			log.Println("Error querying LLM:", err)
			return
		}

		// Handle the output format
		if output_format == "json" {
			results.WriteJSONData(response, filenames[i], outputFile) // Write formatted JSON to file
		} else {
			if output_format == "csv" {
				keys := prompt.GetResultsKeys(config)
				writer := results.CreateCSVWriter(outputFile, keys) // Pass the file to CreateWriter
				defer writer.Flush()                                // Ensure data is flushed after all writes
				results.WriteCSVData(response, filenames[i], writer, keys)
			}
		}

		if config.Project.Configuration.CotJustification == "yes" {
			justificationFilePath := resultsFileName + "_" + filenames[i] + "_justification.txt"
			err := os.WriteFile(justificationFilePath, []byte(justification), 0644)
			if err != nil {
				log.Println("Error writing justification file:", err)
				return
			}
		}

		// Sleep before the next prompt if it's not the last one
		if i < len(prompts)-1 {
			waitWithStatus(getWaitTime(promptText, config))
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
	// Locking to ensure thread-safety when accessing the requestTimestamps slice
	mutex.Lock()
	defer mutex.Unlock()

	// Clean up old timestamps (older than 60 seconds)
	now := time.Now()
	cutoff := now.Add(-60 * time.Second)
	validTimestamps := []time.Time{}
	for _, timestamp := range requestTimestamps {
		if timestamp.After(cutoff) {
			validTimestamps = append(validTimestamps, timestamp)
		}
	}
	requestTimestamps = validTimestamps

	// Add the current request timestamp
	requestTimestamps = append(requestTimestamps, now)

	// Get the current number of requests in the last 60 seconds
	numRequests := len(requestTimestamps)

	// Calculate the time to wait until the next minute
	remainingSeconds := 60 - now.Second()
	// Analyze TPM limits
	tpm_wait_seconds := 0
	// Get the TPM limit from the configuration
	tpmLimit := config.Project.LLM.TpmLimit
	if tpmLimit > 0 {
		// Get the number of tokens from the prompt
		tokens := cost.GetNumTokensFromPrompt(prompt, config)
		tpm_wait_seconds = remainingSeconds
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
			tpm_wait_seconds = remainingSeconds + int(requiredWaitTime) + secondsToMinute
		}
		// Otherwise, calculate the wait time based on tokens used
	}
	// Analyze RPM limits
	rpm_wait_seconds := 0
	rpmLimit := config.Project.LLM.RpmLimit
	if rpmLimit > 0 {
		// If the number of requests risks to exceed the RPM limit, we need to wait
		if numRequests >= rpmLimit-1 {
			rpm_wait_seconds = remainingSeconds
		}
	}

	// Return the maximum of tpm_wait_seconds and rpm_wait_seconds
	if tpm_wait_seconds > rpm_wait_seconds {
		return tpm_wait_seconds
	} else {
		return rpm_wait_seconds
	}
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
