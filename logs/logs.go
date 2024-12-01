package logs

import (
	"log"
	"os"

	"github.com/fatih/color"
)

var Logger *log.Logger
var logToFile bool

// Init initializes the logger based on whether to log to file or not
func Init(logToFileFlag bool) {
	logToFile = logToFileFlag

	// If logging to file, initialize the file logger
	if logToFile {
		file, err := os.OpenFile("argus.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Failed to open log file: %v\n", err)
		}
		Logger = log.New(file, "ARGUS: ", log.LstdFlags)
	} else {
		// If not logging to file, log to console
		Logger = log.New(os.Stdout, "ARGUS: ", log.LstdFlags)
	}
}

// Info logs info messages. If logging to console, prints with color.
func Info(message string) {
	if logToFile {
		Logger.Println("INFO: " + message)
	} else {
		// Print colored info message to console
		color.New(color.FgGreen).Println("INFO: " + message)
	}
}

// Error logs error messages. If logging to console, prints with color.
func Error(err error) {
	if logToFile {
		Logger.Println("ERROR: " + err.Error())
	} else {
		// Print colored error message to console
		color.New(color.FgRed).Println("ERROR: " + err.Error())
	}
}
