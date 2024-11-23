package codec

import (
	"encoding/json"
)

type Json struct {
	schemaVersion string
}

func NewJson(schemaVersion string) *Json {
	return &Json{
		schemaVersion: schemaVersion,
	}
}

func (j Json) Encode(value interface{}) (data []byte, err error) {
	return json.Marshal(value)
}

func (j Json) Decode(data []byte, value interface{}) error {
	return json.Unmarshal(data, value)
}

func (j Json) SchemaVersion() string {
	return j.schemaVersion
}

func (j Json) ContentType() string {
	return "application/json"
}
