package monitor

import (
	"net/http"
	"time"
)

// CheckURL checks the URL and returns the HTTP status code, response time, and any error.
func CheckURL(url string) (int, time.Duration, error) {
	start := time.Now()
	response, err := http.Get(url)
	elapsed := time.Since(start)

	if err != nil {
		return 0, elapsed, err // Return 0 if there is an error, indicating no valid status code
	}

	return response.StatusCode, elapsed, nil
}
