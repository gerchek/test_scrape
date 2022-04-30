package logging

import (
	"fmt"
	"os"
	"time"
)

// getLogFilePath get the log file save path
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", os.Getenv("RuntimeRootPath"), os.Getenv("LogSavePath") /* "runtime/logs/" */)
}

// getLogFileName get the save name of the log file
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		os.Getenv("LogSaveName"),  //"ty_log",
		time.Now().Format(os.Getenv("TimeFormat") /*"02.01.2006"*/),
		os.Getenv("LogFileExt"), //"log",
	)
}
