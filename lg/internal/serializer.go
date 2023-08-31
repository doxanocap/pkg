package internal

import (
	"encoding/json"
)

type LogMsg struct {
	Level   string
	Host    string
	Message string
	Payload string
}

func Marshal(logMsg *LogMsg) string {
	msg, _ := json.Marshal(logMsg)
	return string(msg)
}

func MarshalStruct(v interface{}) string {
	msg, _ := json.Marshal(v)
	return string(msg)
}
