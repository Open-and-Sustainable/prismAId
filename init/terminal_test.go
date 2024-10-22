package init

import (
    "io"
    "os"
    "strings"
    "testing"
)

func TestRunInteractiveConfigCreation(t *testing.T) {
    // Step 1: Prepare the simulated user input
    inputs := strings.Join([]string{
        "./test_config.toml",       // File path to save the configuration
        "Test Project",             // Project name
        "Test Author",              // Author name
        "1.0",                      // Project version
        "./input",                  // Input directory
        "",                         // Input conversion (empty means none)
        "./results",                // Results directory
        "csv",                      // Output format
        "low",                      // Log level
        "no",                       // Duplication
        "no",                       // Chain-of-thought justification
        "no",                       // Document summary
        "no",                       // Do you want to add a model configuration?
        "yes",                      // Do you confirm the standard 'persona' part?
        "yes",                      // Do you confirm the standard 'task' part?
        "yes",                      // Do you confirm the standard 'expected_result' part?
        "yes",                      // Do you confirm the standard 'failsafe' part?
        "no",                       // Do you want to add review item #1?
        // Add any additional inputs as required by your prompts
    }, "\n")

    // Step 2: Redirect os.Stdin
    originalStdin := os.Stdin
    defer func() { os.Stdin = originalStdin }()
    r, w, err := os.Pipe()
    if err != nil {
        t.Fatalf("Failed to create pipe: %v", err)
    }
    os.Stdin = r

    // Write inputs to the pipe
    go func() {
        defer w.Close()
        _, err := io.WriteString(w, inputs)
        if err != nil {
            t.Errorf("Failed to write inputs: %v", err)
        }
    }()

    // Step 3: Run the function
    RunInteractiveConfigCreation()

    // Step 4: Verify the output file was created
    configFilePath := "./test_config.toml"
    if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
        t.Fatalf("Expected configuration file '%s' to be created.", configFilePath)
    }

    // Optionally, read and verify the contents of the generated config file
    content, err := os.ReadFile(configFilePath)
    if err != nil {
        t.Fatalf("Failed to read configuration file: %v", err)
    }

    // Perform assertions on the content
    if !strings.Contains(string(content), `name = "Test Project"`) {
        t.Errorf("Configuration file does not contain expected project name.")
    }

    // Clean up
    os.Remove(configFilePath)
}
