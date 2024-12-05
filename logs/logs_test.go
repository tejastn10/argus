package logs

import (
	"bytes"
	"os"
	"testing"
)

// Helper function to clean up after each test (if necessary)
func cleanup(t *testing.T) {
	// Remove the log file if it exists
	os.Remove("output/argus.log")

	// Remove the output folder if it is empty
	if err := os.RemoveAll("output"); err != nil {
		// Log an error if the folder removal fails
		t.Errorf("Failed to remove output folder: %v", err)
	}
}

// TestLoggingToFile verifies that log messages are correctly logged to a file.
func TestLoggingToFile(t *testing.T) {
	// Clean up from any previous test runs
	cleanup(t)

	// Initialize logger to log to file inside the output folder
	Init(true, false)

	// Log an info message
	Info("File logging test")

	// Check if the log message is written to the output/argus.log file
	fileContent, err := os.ReadFile("output/argus.log")
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	expected := "INFO    : File logging test\n"
	if !contains(fileContent, expected) {
		t.Errorf("Expected log output to contain %q, but got %q", expected, string(fileContent))
	}

	// Clean up
	cleanup(t)
}

// Helper function to check if a byte slice contains a specific string
func contains(content []byte, substr string) bool {
	return bytes.Contains(content, []byte(substr))
}
