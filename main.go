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
	retryCount := flag.Int("retryCount", 3, "The number of retries for monitoring requests. Default is 3 retries")
	backoffDuration := flag.Int("backoffDuration", 2, "The backoff duration (in seconds) between retries. Default is 2 seconds")

	// Parse the flags
	flag.Parse()

	// Initialize logger with a timestamp
	logs.Init(*logToFile, *logTimestamp)

	// Log the start of the monitoring process
	logs.Info(fmt.Sprintf("Starting uptime monitoring for %s every %v seconds", *url, *interval))

	// Monitoring loop
	for {
		status, elapsed, err := monitor.CheckURL(*url)
		if err != nil {
			// Log error with improved structure
			logs.Error(fmt.Errorf("failed to check URL %s | Elapsed Time: %v | Error: %v", *url, elapsed, err))
		} else {
			// Log status and response time using the logs package
			logs.Success(fmt.Sprintf("URL: %s | Response Time: %v | Status: %d", *url, elapsed, status))
		}
		time.Sleep(time.Duration(*interval) * time.Second)
	}
}
