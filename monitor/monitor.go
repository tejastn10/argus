package monitor

import (
	"net/http"
	"time"
)

func CheckURL(url string) (bool, time.Duration, error) {
	start := time.Now()
	response, err := http.Get(url)
	elapsed := time.Since(start)

	if err != nil {
		return false, elapsed, err
	}

	return response.StatusCode == http.StatusOK, elapsed, nil
}
