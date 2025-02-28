package monitor

import (
	"net/http"
	"testing"
	"time"
)

// TestMonitorURLTable tests the MonitorURL function using a table-driven approach.
// Each test case specifies a URL, the expected status code, and whether an error is expected.
func TestMonitorURLTable(t *testing.T) {
	tests := []struct {
		name       string // Name of the test case
		url        string // URL to test
		statusCode int    // Expected HTTP status code
		wantErr    bool   // Whether an error is expected
	}{
		{"Valid URL", "https://example.com", http.StatusOK, false}, // Test a valid URL
		{"Invalid URL", "http://nonexistent-url", 0, true},         // Test an invalid URL
	}

	// Iterate over each test case
	for _, tt := range tests {
		// Run each test case as a sub-test
		t.Run(tt.name, func(t *testing.T) {
			// Call the MonitorURL function with the test case URL
			statusCode, _, err := MonitorURL(tt.url, 3, 2*time.Second)

			// Check if the error behavior matches the expectation
			if (err != nil) != tt.wantErr {
				t.Errorf("MonitorURL() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Verify the returned status code, but only if no error is expected
			if statusCode != tt.statusCode && !tt.wantErr {
				t.Errorf("Expected status code %d, got %d", tt.statusCode, statusCode)
			}
		})
	}
}
