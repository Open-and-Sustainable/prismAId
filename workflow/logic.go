package workflow

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"prismAId/check"
	"prismAId/config"
	"prismAId/convert"
	"prismAId/cost"
	"prismAId/debug"
	"prismAId/model"
	"prismAId/prompt"
	"prismAId/results"
	"prismAId/review"
	"prismAId/tokens"
	"sync"
	"time"
)

const (
	// Define a specific exit code for wrong command call
	ExitCodeWrongCommandCall = 1
	// Define a specific exit code for review logic errors
	ExitCodeErrorInReviewLogic = 2
	// Define a specific exit code for supplier model errors
	ExitCodeInputSupplierModelError = 3
	// Define a specific exit code for input token errors
	ExitCodeInputTokenError = 4
)

// Global variable to store the timestamps of requests
var requestTimestamps []time.Time
var mutex sync.Mutex

// RunReview executes the review process based on the provided configuration file path.
// This function loads the configuration, manages prompts, and coordinates the various modules to complete 
// the review workflow.
//
// Arguments:
// - cfg_path: A string representing the file path to the configuration file.
//
// Returns:
// - An error if there is a problem with the configuration or during the review execution.
func RunReview(cfg_path string) error {
	// load project configuration
	config, err := config.LoadConfig(cfg_path, config.RealFileReader{}, config.RealEnvReader{})
	if err != nil {
		fmt.Println("Error loading project configuration:", err) // here the logging function is not implemented yet
		return err
	}

	// setup logging
	if config.Project.Configuration.LogLevel == "high" {
		debug.SetupLogging(debug.File, cfg_path)
	} else if config.Project.Configuration.LogLevel == "medium" {
		debug.SetupLogging(debug.Stdout, cfg_path)
	} else {
		debug.SetupLogging(debug.Silent, cfg_path) // default value
	}

	// run input conversion if needed
	if config.Project.Configuration.InputConversion != "no" {
		err := convert.Convert(config)
		if err != nil {
			log.Printf("Error:\n%v", err)
			os.Exit(ExitCodeErrorInReviewLogic)
		}
	}

	// setup other debugging features
	if config.Project.Configuration.Duplication == "yes" {
		debug.DuplicateInput(config)
	}

	// generate prompts
	prompts, filenames := prompt.ParsePrompts(config)
	log.Println("Found", len(prompts), "files")

	// build options object
	options, err := review.NewOptions(config.Project.Configuration.ResultsFileName, config.Project.Configuration.OutputFormat, config.Project.Configuration.CotJustification, config.Project.Configuration.Summary)
	if err != nil {
		log.Printf("Error:\n%v", err)
		return err
	}

	// build query object
	query, err := review.NewQuery(prompts, prompt.GetResultsKeys(config))
	if err != nil {
		log.Printf("Error:\n%v", err)
		return err
	}

	// build models object
	models, err := review.NewModels(config.Project.LLM)
	if err != nil {
		log.Printf("Error:\n%v", err)
		return err
	}
	
	// differentiate logic if simgle model review or ensemble
	ensemble := false
	if len(models) > 1 {ensemble = true}

	if ensemble {
		fmt.Println("Cost estimates are available for single model reviews only.")
	}
	
	for _, model := range models {
		if !ensemble {model.ID = ""}
		err = runSingleModelReview(model, options, query, filenames)
		if err != nil {
			log.Printf("Error:\n%v", err)
			return err
		}	
	}
	
	// cleanup eventual debugging temporary files
	if config.Project.Configuration.Duplication == "yes" {
		debug.RemoveDuplicateInput(config)
	}

	log.Println("Done!")
	return nil
}

// Method that returns the number of seconds to wait to respect TPM limits
func getWaitTime(prompt string, llm review.Model) int {
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
	tpmLimit := llm.TPM
	if tpmLimit > 0 {
		// Get the number of tokens from the prompt
		counter := tokens.RealTokenCounter{}
		tokens := counter.GetNumTokensFromPrompt(prompt, llm.Provider, llm.Model, llm.APIKey)
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
	rpmLimit := llm.RPM
	if rpmLimit > 0 {
		// If the number of requests risks to exceed the RPM limit, we need to wait
		if numRequests >= int(rpmLimit-1) {
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
			log.Printf("Waiting... %d seconds remaining\n", remainingTime)
		}
		remainingTime--
		// Break the loop when no time is left
		if remainingTime <= 0 {
			log.Println("Wait completed.")
			break
		}
	}
}

func getDirectoryPath(resultsFileName string) string {
	dir := filepath.Dir(resultsFileName)

	// If the directory is ".", return an empty string
	if dir == "." {
		return ""
	}
	return dir
}

func runSingleModelReview(llm review.Model, options review.Options, query review.Query, filenames []string) error {

	// start writer for results.. the file will be project_name[.csv or .json] in the path where the toml is
	resultsFileName := options.ResultsFileName
	outputFilePath := resultsFileName + "." + options.OutputFormat
	if llm.ID != "" {outputFilePath = resultsFileName + "_" + llm.ID + "." + options.OutputFormat}
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		log.Println("Error creating output file:", err)
		return err
	}
	defer outputFile.Close() // Ensure the file is closed after all operations are done

	var writer *csv.Writer
	if options.OutputFormat == "csv" {
		writer = results.CreateCSVWriter(outputFile, query.Keys) // Pass the file to CreateWriter
		defer writer.Flush()                                // Ensure data is flushed after all writes	
	} else if options.OutputFormat == "json" {
		err = results.StartJSONArray(outputFile)
		if err != nil {
			log.Println("Error starting JSON array:", err)
			return err
		}
	}

	// Loop through the prompts
	for i, promptText := range query.Prompts {
		log.Println("File: ", filenames[i], " Prompt: ", promptText)

		// clean model names
		llm.Model = check.GetModel(promptText, llm.Provider, llm.Model, llm.APIKey)
		fmt.Println("Processing file "+string(i+1)+"/"+string(len(query.Prompts))+" "+filenames[i]+" with model "+llm.Model)
		
		// check if prompts resepct input tokens limits for selected models
		counter := tokens.RealTokenCounter{}
		checkInputLimits := check.RunInputLimitsCheck(promptText, llm.Provider, llm.Model, llm.APIKey, counter)
		if checkInputLimits != nil {
			fmt.Println("Error resepecting the max input tokens limits for the following manuscripts and models.")
			log.Printf("Error:\n%v", checkInputLimits)
			os.Exit(ExitCodeInputTokenError)	
		}
		if llm.ID == "" {			
			// ask if continuing given the total cost
			check := check.RunUserCheck(cost.ComputeCosts(query.Prompts, llm.Provider, llm.Model, llm.APIKey), llm.Provider)
			if check != nil {
				log.Printf("Error:\n%v", check)
				os.Exit(0) // if the user stops the execution it is still a success run, hence exit code = 0, but the reason for the exit may be different hence is logged
			}
		}

		// Query the LLM
		realQueryService := model.DefaultQueryService{}
		response, justification, summary, err := realQueryService.QueryLLM(promptText, llm, options)
		if err != nil {
			log.Println("Error querying LLM:", err)
			return err
		}

		// Handle the output format
		if options.OutputFormat == "json" {
			results.WriteJSONData(response, filenames[i], outputFile) // Write formatted JSON to file
			// add comma if it's not the last element
			if i < len(query.Prompts)-1 {
				results.WriteCommaInJSONArray(outputFile)
			}
		} else {
			if options.OutputFormat == "csv" {
				results.WriteCSVData(response, filenames[i], writer, query.Keys)
			}
		}
		// save justifications
		if options.Justification {
			justificationFilePath := getDirectoryPath(resultsFileName) + "/" + filenames[i] + "_justification.txt"
			if llm.ID != "" {justificationFilePath = getDirectoryPath(resultsFileName) + "/" + filenames[i] + "_justification_"+llm.ID+".txt"}
			err := os.WriteFile(justificationFilePath, []byte(justification), 0644)
			if err != nil {
				log.Println("Error writing justification file:", err)
				return err
			}
		}
		// save summaries
		if options.Summary {
			summaryFilePath := getDirectoryPath(resultsFileName) + "/" + filenames[i] + "_summary.txt"
			if llm.ID != "" {summaryFilePath = getDirectoryPath(resultsFileName) + "/" + filenames[i] + "_summary_"+llm.ID+".txt"}
			err := os.WriteFile(summaryFilePath, []byte(summary), 0644)
			if err != nil {
				log.Println("Error writing summary file:", err)
				return err
			}
		}

		// Sleep before the next prompt if it's not the last one
		if i < len(query.Prompts)-1 {
			waitWithStatus(getWaitTime(promptText, llm))
		}
	}

	// close JSON array if needed
	if options.OutputFormat == "json" {
		err = results.CloseJSONArray(outputFile)
		if err != nil {
			log.Println("Error closing JSON array:", err)
			return err
		}
	}	
	
	return nil
}
