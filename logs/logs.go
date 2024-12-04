package logs

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
)

var Logger *log.Logger
var logToFile bool
var showTimestamp bool // Store the provided timestamp globally

// Init initializes the logger based on whether to log to file or not
// and include timestamp which will be included in all log messages.
func Init(logToFileFlag bool, logTimestamp bool) {
	logToFile = logToFileFlag
	showTimestamp = logTimestamp // Store to show the timestamp

	// If logging to file, initialize the file logger
	if logToFile {
		// Determine the log directory (either local or Docker)
		logDir := "./output" // Default directory, which will be mounted in Docker

		// Ensure the log directory exists
		err := ensureLogDirectory(logDir)
		if err != nil {
			log.Fatalf("Failed to create output directory: %v\n", err)
		}

		// Open the log file inside the output directory
		file, err := os.OpenFile(fmt.Sprintf("%s/argus.log", logDir), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			log.Fatalf("Failed to open log file: %v\n", err)
		}
		Logger = log.New(file, "ARGUS: ", log.LstdFlags)
	} else {
		// If not logging to file, log to console
		Logger = log.New(os.Stdout, "ARGUS: ", log.LstdFlags)
	}
}

// ensureLogDirectory creates the output directory if it doesn't exist
func ensureLogDirectory(logDir string) error {
	// Check if the directory exists
	_, err := os.Stat(logDir)
	if os.IsNotExist(err) {
		// Create the directory if it doesn't exist
		err = os.MkdirAll(logDir, 0750)
	}
	return err
}

// LogMessage formats the log message with colored components and prints it
func LogMessage(level, message string) {
	// Get timestamp with yellow color (only for console)
	timestamp := getTimestamp() // For file logging, we don't apply color to timestamp

	// Format the level with the appropriate color (only for console)
	var levelColor *color.Color
	switch level {
	case "INFO":
		levelColor = color.New(color.FgBlue)
	case "SUCCESS":
		levelColor = color.New(color.FgGreen)
	case "WARNING":
		levelColor = color.New(color.FgYellow)
	case "ERROR":
		levelColor = color.New(color.FgRed)
	default:
		levelColor = color.New(color.FgWhite)
	}

	// Padding for the log level (for consistency in formatting)
	paddedLevel := fmt.Sprintf("%-8s", level) // 8 spaces for padding

	// Print the formatted log message
	var logMessage string
	if logToFile {
		// Format the log message with the padded level
		logMessage = paddedLevel + ": " + message
		// Log to file without colors
		Logger.Println(logMessage) // This logs the message without any color
	} else {
		// Format the log message with the padded level
		if showTimestamp {
			logMessage = color.New(color.FgYellow).Sprint(timestamp) + " " + levelColor.Sprint(paddedLevel) + ": " + message
		} else {
			logMessage = levelColor.Sprint(paddedLevel) + ": " + message
		}
		// Print to console with color formatting
		_, err := color.New(color.Reset).Println(logMessage)
		if err != nil {
			log.Fatalf("Failed to print log message: %v", err)
		}
	}
}

// Info logs info messages
func Info(message string) {
	LogMessage("INFO", message)
}

// Success logs success messages
func Success(message string) {
	LogMessage("SUCCESS", message)
}

// Warning logs warning messages
func Warning(message string) {
	LogMessage("WARNING", message)
}

// Error logs error messages
func Error(err error) {
	LogMessage("ERROR", err.Error())
}

// getTimestamp returns the current timestamp as a string in the format YYYY-MM-DD HH:MM:SS
func getTimestamp() string {
	// Get the current time
	currentTime := time.Now()

	// Format the time as YYYY-MM-DD HH:MM:SS
	return currentTime.Format("2006-01-02 15:04:05")
}
