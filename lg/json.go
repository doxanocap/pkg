package lg

import (
	"encoding/json"
)

type LogMsg struct {
	Time    string `json:"time"`
	Level   string `json:"level"`
	Host    string `json:"host"`
	File    string `json:"file"`
	Line    int    `json:"line"`
	Message string `json:"message"`
}

func Marshal(logMsg *LogMsg) string {
	msg, _ := json.Marshal(logMsg)
	return string(msg)
}

func MarshalStruct(v interface{}) string {
	msg, _ := json.Marshal(v)
	return string(msg)
}
