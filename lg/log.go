package lg

import (
	"fmt"
	"github.com/doxanocap/pkg/lg/internal"
	"log"
	"os"
)

type mode bool

// Global
var (
	Mode mode = Regular
	l         = log.New(os.Stderr, "", log.LstdFlags)
)

// Global
const (
	Tracing mode = true
	Regular mode = false
)

// Log with log level
func Log(level, msg string) {
	l.Println(internal.Marshal(&internal.LogMsg{
		Level:   level,
		Host:    internal.GetHost(),
		Message: msg,
		Payload: msg,
	}))
}

// LogInfo for positive events
func LogInfo(msg string) {
	l.Println(internal.Marshal(&internal.LogMsg{
		Level:   "INFO",
		Host:    internal.GetHost(),
		Message: msg,
		Payload: msg,
	}))
}

// LogTrace for trace data
func LogTrace(msg string, spot interface{}) {
	if Mode {
		spotMsg := ""
		switch spot.(type) {
		case interface{}:
			spotMsg = internal.MarshalStruct(spot)
		case int:
			spotMsg = fmt.Sprintf("%+d", spot)
		default:
			spotMsg = fmt.Sprintf("%s", spot)
		}

		l.Println(internal.Marshal(&internal.LogMsg{
			Level:   "TRACE",
			Host:    internal.GetHost(),
			Message: msg,
			Payload: msg + ", " + spotMsg,
		}))
	}
}

// LogFatal stop app use os.Exit(1)
func LogFatal(err error) {
	errStr := cast(err)
	l.Fatalln(internal.Marshal(&internal.LogMsg{
		Level:   "FATAL",
		Host:    internal.GetHost(),
		Message: err.Error(),
		Payload: errStr,
	}))
}

// LogError print error with call data
func LogError(err error) {
	errStr := cast(err)

	call := internal.CallInfo(2)
	l.Println(internal.Marshal(&internal.LogMsg{
		Level:   "ERROR",
		Host:    internal.GetHost(),
		Message: err.Error(),
		Payload: errStr + ", " + internal.MarshalStruct(call),
	}))
}

func cast(err error) string {
	return err.Error()
}
