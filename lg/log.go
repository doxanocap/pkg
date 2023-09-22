package lg

import (
	"fmt"
	"github.com/doxanocap/pkg/errs"
	"log"
	"os"
	"time"
)

const (
	levelINFO  = "INFO"
	levelERROR = "ERROR"
	levelWARN  = "WARN"
	levelFATAL = "FATAL"

	callLevel = 3
)

// Global
var (
	timeFormat = time.RFC3339
	global     = log.New(os.Stdout, "", log.Lmsgprefix)
)

// Log Logging with custom log level
func Log(level, msg string) {
	encoder := newEncoder()
	call := CallInfo(callLevel)

	fmt.Println(encoder.Marshal(&LogMsg{
		Time:    time.Now().Format(timeFormat),
		Level:   level,
		Host:    GetHost(),
		File:    call.FileName,
		Line:    call.Line,
		Message: msg,
	}))

	stringifierLogMsg := encoder.Marshal(&LogMsg{
		Time:    time.Now().Format(timeFormat),
		Level:   level,
		Host:    GetHost(),
		File:    call.FileName,
		Line:    call.Line,
		Message: msg,
	})

	if level == levelFATAL {
		global.Fatalln(stringifierLogMsg)
		return
	}

	global.Println(stringifierLogMsg)
}

// Info Logging for positive events
func Info(msg string) {
	Log(levelINFO, msg)
}

// Infof Logging for positive events with formatting
func Infof(format string, input ...any) {
	Log(levelINFO, fmt.Sprintf(format, input))
}

// Warn Logging warnings
func Warn(msg string) {
	Log(levelWARN, msg)
}

// Warnf Logging warnings with formatting
func Warnf(format string, input ...any) {
	Log(levelWARN, fmt.Sprintf(format, input...))
}

// Fatal Logging with newline and stops app using os.Exit(1)
func Fatal(v ...any) {
	Log(levelFATAL, fmt.Sprint(v...))
}

// Fatalf Logging with newline, formatting, and stops app using os.Exit(1)
func Fatalf(format string, input ...any) {
	Log(levelFATAL, fmt.Sprintf(format, input...))
}

// Error Logging error with call data
func Error(err error) {
	msg := "nil"
	if err != nil {
		msg = err.Error()
	}
	Log(levelERROR, msg)
}

// Errorf Logging error with call data and formatting
func Errorf(format string, input ...any) {
	Log(levelERROR, fmt.Sprintf(format, input))
}

// LError Logging CustomError with levels and methods DEPRECATED
func LError(err errs.CustomError) {
	//errStr := err.Error()
	//call := internal.CallInfo(2)
	//global.Println(internal.Marshal(&internal.LogMsg{
	//	Level:   levelERROR,
	//	Host:    internal.GetHost(),
	//	Message: errStr,
	//}))
}

func SetTimeFormat(customFormat string) {
	timeFormat = customFormat
}
