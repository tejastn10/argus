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

	// Parse the flags
	flag.Parse()

	// Initialize logger with a timestamp
	logs.Init(*logToFile, *logTimestamp)

	// URL to monitor and monitoring interval
	url := "https://example.com"
	interval := 30 * time.Second

	// Log the start of the monitoring process
	logs.Info(fmt.Sprintf("Starting uptime monitoring for %s every %v", url, interval))

	// Monitoring loop
	for {
		status, elapsed, err := monitor.CheckURL(url)
		if err != nil {
			// Log error with improved structure
			logs.Error(fmt.Errorf("failed to check URL %s | Elapsed Time: %v | Error: %v", url, elapsed, err))
		} else {
			// Log status and response time using the logs package
			logs.Success(fmt.Sprintf("URL: %s | Response Time: %v | Status: %d", url, elapsed, status))
		}
		time.Sleep(interval)
	}
}
