/*
    This logger is used to log data and cache this logs,
	into JSON file, so we can use it in our frontend as well as debugging our backend.
	---------------
	Instructions to use this lib

	first you should use [logger.SetupLoggingSystem()] in the main function,
	so that you configure the logger
*/

package logger

import (
	"encoding/json"
	"fmt"
	"github.com/TwiN/go-color"
	"log"
	"os"
	"time"
)

// Fields:
var logFileNameWithPath = "./logs.json"

// LogType for using it as enum:
type LogType int64

const (
	ERROR       LogType = 0
	WARNING     LogType = 1
	INFORMATION LogType = 2
	SUCCESS     LogType = 3
)

func (logType LogType) String() string {
	// Initializing:
	name := "ERROR"
	// Checking:
	switch logType {
	// Cases:
	case ERROR:
		name = "ERROR"
	case INFORMATION:
		name = "INFORMATION"
	case WARNING:
		name = "WARNING"
	case SUCCESS:
		name = "SUCCESS"
	}
	// Returning:
	return name
}

// LogEntity Structs:
type LogEntity struct {
	// Fields:
	Id      int64  `json:"id"`
	Tag     string `json:"tag"`
	Message string `json:"message"`
	// FullMessage  string  `json:"full_message"`
	Type         LogType `json:"type"`
	TypeAsString string  `json:"TypeAsString"`
	Date         string  `json:"date"`       // As human-readable string format
	Timestamps   int64   `json:"timestamps"` // As timestamps
}

/*
SetupLoggingSystem this method is responsible for setup and initialize the logger system,
you can set a default path for storing the logs file, or you can pass empty string to use the default
path which is "./logs.json".
*/
func SetupLoggingSystem(customLogFileNameWithPath string) {
	// Checking: if dev entered a custom path.
	if customLogFileNameWithPath != "" {
		// Changing:
		logFileNameWithPath = customLogFileNameWithPath
	}
	// Checking:
	if _, err := os.Stat(logFileNameWithPath); err != nil {
		/// File is not exits,
		/// so we'll create new one & write the initial JSON array.

		// Initializing:
		file, err := os.Create(logFileNameWithPath)

		// Checking:
		if err != nil {
			// log.Fatal & Print the error.
			log.Fatal(err)
		}

		// Writing: the initial JSON array.
		_, writerErr := file.WriteString("[]")

		// Checking:
		if writerErr != nil {
			// log.Fatal & Print the error.
			log.Fatal(writerErr)
		}

		// Closing:
		defer func(file *os.File) {
			// Closing: current file.
			err := file.Close()
			// Checking:
			if err != nil {
				// log.Fatal & Print the error.
				log.Fatal(err)
			}
		}(file)
	}
}

func readLogs() []LogEntity {
	// Reading:
	file, err := os.ReadFile(logFileNameWithPath)
	var logs []LogEntity

	// Reading:
	err2 := json.Unmarshal(file, &logs)

	// Checking:
	if err != nil || err2 != nil {
		// log.Fatal & Print the error.
		log.Fatal(err)
	}

	// Returning:
	return logs
}

func displayLogWithStyle(
	// Parameters:
	date string,
	tag string,
	message string,
	logType LogType,
) {
	// Checking:
	dateColor := color.GrayBackground
	typeColor := color.RedBackground
	tagsColor := color.CyanBackground
	// Checking:
	switch logType {
	// Cases:
	case ERROR:
		typeColor = color.RedBackground
	case INFORMATION:
		typeColor = color.BlueBackground
	case WARNING:
		typeColor = color.YellowBackground
	case SUCCESS:
		typeColor = color.GreenBackground
	}
	// Displaying:
	fmt.Printf(
		color.Colorize(dateColor, " %s  ")+
			color.Colorize(typeColor, " %s ")+
			color.Colorize(tagsColor, " %s: ")+
			" "+
			"%s\n",
		date, logType.String(), tag, message,
	)
}

func logData(tag string, message string, logType LogType) {
	// Initializing:
	date := time.Now().Format("2006-01-02 3:4:5 pm")
	timestamps := time.Now().UnixMilli()
	// Logging:
	displayLogWithStyle(date, tag, message, logType)
	// Checking:
	if len(tag) < 3 {
		// LOGGING:
		logData("AR-Logger", "Tags must be at least 3 len for better formatting!", WARNING)
	}
	// Reading:
	logs := readLogs()
	newLog := LogEntity{
		Id:           int64(len(logs) + 1),
		Message:      message,
		Type:         logType,
		TypeAsString: logType.String(),
		Tag:          tag,
		Date:         date,
		Timestamps:   timestamps,
	}

	// Appending:
	logs = append(logs, newLog)

	// Converting: the logs into JSON.
	logsAsJson, err2 := json.Marshal(logs)

	// Writing:
	err := os.WriteFile(logFileNameWithPath, logsAsJson, 0755)

	// Checking:
	if err != nil || err2 != nil {
		// log.Fatal & Print the error.
		log.Fatal(err)
	}
}

// Logs methods:

func E(tag string, error string) {
	// Logging:
	logData(tag, error, ERROR)
}
func W(tag string, warning string) {
	// Logging:
	logData(tag, warning, WARNING)
}

func I(tag string, information string) {
	// Logging:
	logData(tag, information, INFORMATION)
}

func S(tag string, success string) {
	// Logging:
	logData(tag, success, SUCCESS)
}
