package main

import (
	"fmt"
	"time"

	"github.com/tejastn10/argus/monitor"
)

func main() {
	url := "https://example.com"
	interval := 30 * time.Second

	fmt.Printf("Starting uptime monitoring for %s every %v\n", url, interval)
	for {
		status, elapsed, err := monitor.CheckURL(url)
		if err != nil {
			fmt.Printf("Error checking URL %s: %v\n", url, err)
		} else {
			fmt.Printf("URL: %s | Status: %t | Response Time: %v\n", url, status, elapsed)
		}
		time.Sleep(interval)
	}
}
