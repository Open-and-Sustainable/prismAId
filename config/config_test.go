package config

import (
	"os"
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	t.Run("CorrectConfig", func(t *testing.T) {
		// Setup
		expected := Config{APIKey: "12345", Endpoint: "http://api.example.com"}
		os.WriteFile("test.toml", []byte("APIKey = '12345'\nEndpoint = 'http://api.example.com'"), 0644)
		defer os.Remove("test.toml")

		// Execution
		config, err := LoadConfig("test.toml")
		if err != nil {
			t.Fatal("Failed to load config", err)
		}

		// Assertion
		if !reflect.DeepEqual(config, expected) {
			t.Errorf("Expected %v, got %v", expected, config)
		}
	})

	t.Run("MissingFile", func(t *testing.T) {
		_, err := LoadConfig("nonexistent.toml")
		if err == nil {
			t.Error("Expected error for missing file, got none")
		}
	})
}
