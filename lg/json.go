package lg

import (
	"bytes"
	"encoding/json"
)

var (
	jsonEncoder *json.Encoder
)

type Encoder struct {
	buff []byte
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
	var buff []byte

	if jsonEncoder == nil {
		writer := bytes.NewBuffer(buff)
		jsonEncoder = json.NewEncoder(writer)
		jsonEncoder.SetEscapeHTML(false)
	}
	return &Encoder{
		json: jsonEncoder,
		buff: buff,
	}
}

func (e *Encoder) Marshal(logMsg *LogMsg) string {
	_ = e.json.Encode(logMsg)
	e.json..
		fmt.Println(e.buff)
	return string(e.buff)
}

func (e *Encoder) MarshalStruct(v interface{}) string {
	_ = e.json.Encode(v)
	return string(e.buff)
}
