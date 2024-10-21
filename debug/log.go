package debug

import (
	"io"
	"log"
	"os"
	"strings"
)


type LogLevel int
const (
	Silent LogLevel = iota
	Stdout
	File
)

// SetupLogging configures the logging output based on the specified log level. The function supports multiple 
// output destinations, including silent (no logging), standard output, and file logging.
//
// Arguments:
// - level: A LogLevel value that determines the output destination. It can be Silent, Stdout, or File.
// - filename: The name of the file where logs should be stored if the File log level is selected.
//
// This function creates or appends to a log file with the same name as the specified filename (excluding its
// extension) but with a ".log" extension.
func SetupLogging(level LogLevel, filename string) {
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