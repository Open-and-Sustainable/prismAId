package debug

import (
    "bytes"
    "io"
    "log"
    "os"
	"strings"
    "sync"
    "testing"
)

func TestSetupLogging_Silent(t *testing.T) {
    // Mutex to prevent race conditions due to global state modification
    var mu sync.Mutex
    mu.Lock()
    defer mu.Unlock()

    // Save original logger settings
    originalOutput := log.Writer()
    originalFlags := log.Flags()
    defer func() {
        log.SetOutput(originalOutput)
        log.SetFlags(originalFlags)
    }()

    // Create a buffer to capture logs
    var buf bytes.Buffer
    log.SetOutput(&buf)

    // Call SetupLogging with Silent level
    SetupLogging(Silent, "test.log")

    // Log something
    log.Println("This is a test log.")

    // Check that nothing was logged
    if buf.Len() != 0 {
        t.Errorf("Expected no logs, but got: %s", buf.String())
    }
}

func TestSetupLogging_Stdout(t *testing.T) {
    var mu sync.Mutex
    mu.Lock()
    defer mu.Unlock()

    // Save original logger settings
    originalOutput := log.Writer()
    originalFlags := log.Flags()
    originalStdout := os.Stdout
    defer func() {
        log.SetOutput(originalOutput)
        log.SetFlags(originalFlags)
        os.Stdout = originalStdout
    }()

    // Create a pipe to capture stdout
    r, w, err := os.Pipe()
    if err != nil {
        t.Fatalf("Failed to create pipe: %v", err)
    }
    os.Stdout = w

    // Call SetupLogging with Stdout level
    SetupLogging(Stdout, "test.log")

    // Log something
    log.Println("This is a test log.")

    // Close writer to flush
    w.Close()

    // Read captured output
    var buf bytes.Buffer
    io.Copy(&buf, r)

    // Check that the log was written to stdout
    if !bytes.Contains(buf.Bytes(), []byte("This is a test log.")) {
        t.Errorf("Expected log output, but got: %s", buf.String())
    }
}

func TestSetupLogging_File(t *testing.T) {
    var mu sync.Mutex
    mu.Lock()
    defer mu.Unlock()

    // Save original logger settings
    originalOutput := log.Writer()
    originalFlags := log.Flags()
    defer func() {
        log.SetOutput(originalOutput)
        log.SetFlags(originalFlags)
    }()

    // Create a temporary file
    tmpfile, err := os.CreateTemp("", "test*.toml")
    if err != nil {
        t.Fatal(err)
    }
    defer os.Remove(tmpfile.Name()) // Clean up

    // Determine the log file name
    logname := strings.TrimSuffix(tmpfile.Name(), ".toml") + ".log"

    // Ensure the log file is removed before the test
    os.Remove(logname)
    defer os.Remove(logname) // Clean up after test

    // Call SetupLogging with File level
    SetupLogging(File, tmpfile.Name())

    // Log something
    log.Println("This is a test log.")

    // Read the log file contents
    content, err := os.ReadFile(logname)
    if err != nil {
        t.Fatalf("Failed to read log file: %v", err)
    }

    // Check that the log was written to the file
    if !strings.Contains(string(content), "This is a test log.") {
        t.Errorf("Expected log output in file, but got: %s", string(content))
    }
}
