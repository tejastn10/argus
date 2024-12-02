package logs

import (
	"bytes"
	"os"
	"testing"
)

// Helper function to clean up after each test (if necessary)
func cleanup() {
	os.Remove("argus.log") // Delete the argus.log file after tests to avoid issues in subsequent runs
}

// TestLoggingToFile verifies that log messages are correctly logged to a file.
func TestLoggingToFile(t *testing.T) {
	// Clean up from any previous test runs
	cleanup()

	// Initialize logger to log to file
	Init(true, false)

	// Log an info message
	Info("File logging test")

	// Check if the log message is written to the argus.log file
	fileContent, err := os.ReadFile("argus.log")
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	expected := "INFO    : File logging test\n"
	if !contains(fileContent, expected) {
		t.Errorf("Expected log output to contain %q, but got %q", expected, string(fileContent))
	}

	// Clean up
	cleanup()
}

// Helper function to check if a byte slice contains a specific string
func contains(content []byte, substr string) bool {
	return bytes.Contains(content, []byte(substr))
}
