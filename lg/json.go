package lg

import (
	"bytes"
	"encoding/json"
	"io"
)

var baseEncoder = &Encoder{}

type Encoder struct {
	buff *bytes.Buffer
	json *json.Encoder
}

type LogMsg struct {
	Time    string `json:"time"`
	Level   string `json:"level"`
	Host    string `json:"host"`
	File    string `json:"file"`
	Line    int    `json:"line"`
	Message string `json:"message"`
}

func newEncoder() *Encoder {
	if baseEncoder.json == nil {
		buff := bytes.NewBuffer([]byte{})
		baseEncoder.buff = buff
		baseEncoder.json = json.NewEncoder(baseEncoder.buff)
		baseEncoder.json.SetEscapeHTML(false)
	}

	return baseEncoder
}

func (e *Encoder) Marshal(logMsg *LogMsg) string {
	_ = e.json.Encode(logMsg)
	res, _ := io.ReadAll(e.buff)
	return string(res)
}

func (e *Encoder) MarshalStruct(v interface{}) string {
	_ = e.json.Encode(v)
	res, _ := io.ReadAll(e.buff)
	return string(res)
}
