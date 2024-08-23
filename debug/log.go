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

// Setup logging based on log level
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