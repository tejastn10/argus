package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/tejastn10/argus/logs"
	"github.com/tejastn10/argus/monitor"
)

func main() {
	// Define the flag for logging behavior
	logToFile := flag.Bool("logToFile", false, "Set to true to log to file, false to log to console")
	logTimestamp := flag.Bool("logTimeStamp", true, "Set to true to log timestamps. Timestamps are logged by default in the file")
	url := flag.String("url", "https://example.com", "The URL to monitor. Default is https://example.com")
	interval := flag.Int("interval", 30, "The monitoring interval in seconds. Default is 30 seconds")
	retryCount := flag.Int("retryCount", 3, "The number of retries for monitoring requests. Must be >= 3. Default is 3 retries")
	backoffDuration := flag.Int("backoffDuration", 3, "The backoff duration (in seconds) between retries. Must be >= 3. Default is 3 seconds")

	// Parse the flags
	flag.Parse()

	// Validate and enforce minimum values for retryCount and backoffDuration
	if *retryCount < 3 {
		logs.Warning("Invalid retryCount provided. Enforcing minimum value of 3.")
		*retryCount = 3
	}
	if *backoffDuration < 3 {
		logs.Warning("Invalid backoffDuration provided. Enforcing minimum value of 3 seconds.")
		*backoffDuration = 3
	}

	// Initialize logger with a timestamp
	logs.Init(*logToFile, *logTimestamp)

	// Log the start of the monitoring process
	logs.Info(fmt.Sprintf("Starting uptime monitoring for %s every %v seconds with %d retries and %d seconds backoff duration",
		*url, *interval, *retryCount, *backoffDuration))

	// Variable to track consecutive failures
	failureCount := 0

	// Monitoring loop
	for {
		// Convert backoffDuration to time.Duration
		backoff := time.Duration(*backoffDuration) * time.Second

		// Call MonitorURL with retryCount and backoffDuration
		status, elapsed, err := monitor.MonitorURL(*url, *retryCount, backoff)
		if err != nil {
			// Increment failure count
			failureCount++

			// Log error with improved structure
			logs.Error(fmt.Errorf("URL: %s | Elapsed Time: %v | Error: %v | Failure Count: %d", *url, elapsed, err, failureCount))

			// If 5 consecutive failures occur, print the message and break the loop
			if failureCount >= 5 {
				logs.Warning(fmt.Sprintf("Service is down: 5 consecutive failures detected for URL %s. Please check your server logs and monitoring systems for further details.", *url))
				logs.Fatal("Stopping monitoring due to 5 consecutive failures.")
				break
			}
		} else {
			// Log status and response time using the logs package
			logs.Success(fmt.Sprintf("URL: %s | Response Time: %v | Status: %d", *url, elapsed, status))

			// Reset failure count on success
			failureCount = 0
		}

		// Wait for the next monitoring interval
		time.Sleep(time.Duration(*interval) * time.Second)
	}
}
