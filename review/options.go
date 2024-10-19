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

func NewOptions(resultsFilename string, outputFormat string, justification string, summary string) (*Options, error) {
	boolJustification := false
	if justification == "yes" {boolJustification = true}
	boolSummary := false
	if summary == "yes" {boolSummary = true}

	return &Options{
		ResultsFileName: resultsFilename,
		OutputFormat:    outputFormat,
		Justification:   boolJustification,
		Summary:      	 boolSummary,
	}, nil
}
