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

	// Parse the flags
	flag.Parse()

	// Initialize the logger based on the flag value
	logs.Init(*logToFile)

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
			logs.Error(fmt.Errorf("failed to check URL %s: %v (elapsed time: %v)", url, err, elapsed))
		} else {
			// Log status and response time using the logs package
			logs.Info(fmt.Sprintf("URL: %s | Status: %d | Response Time: %v", url, status, elapsed))
		}
		time.Sleep(interval)
	}
}
