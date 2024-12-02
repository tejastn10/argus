package monitor

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// CheckURL checks the URL and returns the HTTP status code, response time, and any error.
// It enforces HTTPS and implements a timeout.
func CheckURL(urlString string) (int, time.Duration, error) {
	// Ensure the URL is valid and parse it
	parsedURL, err := url.Parse(urlString)
	if err != nil || !parsedURL.IsAbs() {
		return 0, 0, errors.New("invalid or non-absolute URL")
	}

	// Enforce that the URL uses HTTPS
	if parsedURL.Scheme != "https" {
		return 0, 0, errors.New("only HTTPS URLs are allowed")
	}

	// Check for common URL schemes or patterns to mitigate SSRF
	if strings.Contains(parsedURL.Host, "localhost") || strings.Contains(parsedURL.Host, "127.0.0.1") {
		return 0, 0, errors.New("localhost or private IPs are not allowed")
	}

	// Set a reasonable timeout for the HTTP request to mitigate hanging issues
	client := &http.Client{
		Timeout: 10 * time.Second, // 10 second timeout
	}

	start := time.Now()
	response, err := client.Get(urlString)
	elapsed := time.Since(start)

	if err != nil {
		return 0, elapsed, err // Return 0 if there is an error, indicating no valid status code
	}

	// Ensure the connection was successful (2xx status codes)
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return response.StatusCode, elapsed, errors.New("non-success status code")
	}

	return response.StatusCode, elapsed, nil
}
