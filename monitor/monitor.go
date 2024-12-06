package monitor

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/tejastn10/argus/logs"
)

// MonitorURL checks the URL and returns the HTTP status code, response time, and any error.
// It enforces HTTPS and implements a timeout with retry logic.
// If retryCount is not provided or is 0, it defaults to 3 retries.
func MonitorURL(urlString string, retryCount int, initialBackoff time.Duration) (int, time.Duration, error) {
	// Set default retry count if not provided
	if retryCount == 0 {
		retryCount = 3
	}

	// Set default backoff duration if not provided or zero
	if initialBackoff <= 0 {
		logs.Warning("Invalid backoffDuration provided. Enforcing minimum value of 3 seconds.")
		initialBackoff = 3 * time.Second
	}

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

	var lastError error
	var response *http.Response
	var elapsed time.Duration

	// Retry logic with exponential backoff
	for attempt := 1; attempt <= retryCount; attempt++ {
		start := time.Now()
		response, lastError = client.Get(urlString)
		elapsed = time.Since(start)

		// If the request was successful and returned a valid status code
		if lastError == nil && response.StatusCode >= 200 && response.StatusCode < 300 {
			return response.StatusCode, elapsed, nil
		}

		// If the request failed, log the error and retry after backoff
		if lastError != nil {
			lastError = errors.New("error during request: " + lastError.Error())
		} else {
			lastError = errors.New("non-success status code: " + http.StatusText(response.StatusCode))
		}

		// If the response was received but the status code was not successful, return the status code immediately
		if response != nil && (response.StatusCode < 200 || response.StatusCode >= 300) {
			return response.StatusCode, elapsed, lastError
		}

		// Exponential backoff logic
		backoffDuration := initialBackoff * time.Duration(1<<uint(attempt-1)) // Exponential backoff
		logs.Warning(fmt.Sprintf("Backoff before retrying... Attempt %d of %d. Waiting for %v", attempt, retryCount, backoffDuration))
		time.Sleep(backoffDuration)
	}

	// If all retries failed, return the last error encountered and the response status code if available
	if response != nil {
		return response.StatusCode, elapsed, lastError
	}

	return 0, elapsed, lastError
}
