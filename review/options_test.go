package review

import (
    "testing"
)

func TestNewOptions(t *testing.T) {
    tests := []struct {
        name            string
        resultsFilename string
        outputFormat    string
        justification   string
        summary         string
        want            Options
        wantErr         bool
    }{
        {
            name:            "All positive options",
            resultsFilename: "result.csv",
            outputFormat:    "csv",
            justification:   "yes",
            summary:         "yes",
            want: Options{
                ResultsFileName: "result.csv",
                OutputFormat:    "csv",
                Justification:   true,
                Summary:         true,
            },
            wantErr: false,
        },
        {
            name:            "All negative options",
            resultsFilename: "result.json",
            outputFormat:    "json",
            justification:   "no",
            summary:         "no",
            want: Options{
                ResultsFileName: "result.json",
                OutputFormat:    "json",
                Justification:   false,
                Summary:         false,
            },
            wantErr: false,
        },
        {
            name:            "Mixed options",
            resultsFilename: "output.txt",
            outputFormat:    "txt",
            justification:   "yes",
            summary:         "no",
            want: Options{
                ResultsFileName: "output.txt",
                OutputFormat:    "txt",
                Justification:   true,
                Summary:         false,
            },
            wantErr: false,
        },
        // Additional cases can be added to handle edge cases
        // like incorrect input values if your function should handle these
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := NewOptions(tt.resultsFilename, tt.outputFormat, tt.justification, tt.summary)
            if (err != nil) != tt.wantErr {
                t.Errorf("NewOptions() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("NewOptions() got = %v, want %v", got, tt.want)
            }
        })
    }
}
