package review

import (
)

// options defines the options of the review project
type Options struct {
	ResultsFileName string
	OutputFormat    string
	Justification   bool
	Summary      	bool
}

// NewOptions creates and returns an Options instance based on the provided parameters.
// It interprets certain string inputs as boolean values for the Justification and Summary fields.
//
// Arguments:
// - resultsFilename: A string specifying the name of the file to store the results.
// - outputFormat: A string specifying the format of the output (e.g., "csv", "json").
// - justification: A string that should be "yes" or "no" to determine if justifications are included.
// - summary: A string that should be "yes" or "no" to determine if summaries are included.
//
// Returns:
// - An Options instance with the specified settings.
// - An error if the creation fails, although the current implementation does not anticipate errors.
func NewOptions(resultsFilename string, outputFormat string, justification string, summary string) (Options, error) {
	boolJustification := false
	if justification == "yes" {boolJustification = true}
	boolSummary := false
	if summary == "yes" {boolSummary = true}

	return Options{
		ResultsFileName: resultsFilename,
		OutputFormat:    outputFormat,
		Justification:   boolJustification,
		Summary:      	 boolSummary,
	}, nil
}
