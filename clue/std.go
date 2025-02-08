package clue

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type std struct {
	Code    string `json:"responseCode"`
	Message string `json:"responseMessage"`
	*Clue
}

// GetCode implements Meta.
func (s *std) GetCode() string {
	return s.Code
}

// GetInfo implements Meta.
func (s *std) GetInfo() interface{} {
	return nil
}

// GetMessage implements Meta.
func (s *std) GetMessage() string {
	return s.Message
}

// Templating implements Meta.
func (s *std) Templating(ctx context.Context, clue *Clue) *Clue {
	//modified code
	var (
		code    int    = clue.HttpCode
		casee   string = clue.Meta.GetCode()
		message string = clue.Meta.GetMessage()
	)

	if casee == "" {
		casee = "00"
	}

	if message == "" && code == http.StatusOK {
		message = "Successful"
	}

	clue.Meta.SetCode(fmt.Sprintf("%d%s", code, casee))
	clue.Meta.SetMessage(message)

	return clue
}

// Marshal implements Meta.
func (s *std) MarshalJSON() ([]byte, error) {
	first, err := json.Marshal(s.Clue)
	if err != nil {
		return nil, err
	}
	data := make(map[string]interface{})
	err = json.Unmarshal(first, &data)
	if err != nil {
		return nil, err
	}
	if s.Clue.Meta != nil {
		second, err := json.Marshal(s.Clue.Meta)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(second, &data)
		if err != nil {
			return nil, err
		}
	}
	data["data"] = nil
	if s.Clue.Data != nil {
		var fieldData interface{}
		d, err := json.Marshal(s.Clue.Data)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(d, &fieldData)
		if err != nil {
			return nil, err
		}
		data["data"] = fieldData
	}
	return json.Marshal(data)
}

// SetCode implements Meta.
func (s *std) SetCode(v string) {
	s.Code = v
}

// SetMessage implements Meta.
func (s *std) SetMessage(v string) {
	s.Code = v
}

func MewStd(code, message string) Meta {
	return &std{
		Code:    code,
		Message: message,
	}
}
