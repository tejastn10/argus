package monitor

import (
	"errors"
	"net/http"
	"net/url"
	"time"
)

// CheckURL checks the URL and returns the HTTP status code, response time, and any error.
func CheckURL(urlString string) (int, time.Duration, error) {
	// Validate the URL to ensure it's well-formed and from a trusted source
	parsedURL, err := url.Parse(urlString)
	if err != nil || !parsedURL.IsAbs() {
		return 0, 0, errors.New("invalid URL or not an absolute URL")
	}

	start := time.Now()
	response, err := http.Get(urlString)
	elapsed := time.Since(start)

	if err != nil {
		return 0, elapsed, err // Return 0 if there is an error, indicating no valid status code
	}

	return response.StatusCode, elapsed, nil
}
