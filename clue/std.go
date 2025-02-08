package clue

import (
	"context"
	"encoding/json"
	"net/http"
)

type (
	Pagination struct {
		Page  int `query:"page" json:"page" bson:"page"`
		Limit int `query:"limit" json:"limit" bson:"limit"`
	}
	Info struct {
		Count     int64 `json:"count" bson:"count"`
		TotalPage int64 `json:"total_page" bson:"total_page"`
	}
)

type std struct {
	Code    string `json:"status"`
	Message string `json:"message"`
	Info    *Info  `json:"info,omitempty"`
}

// SetInfo implements Meta.
func (s *std) SetInfo(v *Info) {
	s.Info = v
}

// GetCode implements Meta.
func (s *std) GetCode() string {
	return s.Code
}

// GetInfo implements Meta.
func (s *std) GetInfo() *Info {
	return s.Info
}

// GetMessage implements Meta.
func (s *std) GetMessage() string {
	return s.Message
}

// Templating implements Meta.
func (s *std) Templating(ctx context.Context, clue *Clue) *Clue {
	//modified code
	var (
		code    string = clue.Meta.GetCode()
		message string = clue.Meta.GetMessage()
		info    *Info  = clue.Meta.GetInfo()
	)

	if message == "" && clue.HttpCode == http.StatusOK {
		message = "Successful"
	}

	clue.Meta.SetCode(code)
	clue.Meta.SetMessage(message)
	clue.Meta.SetInfo(info)

	return clue
}

// Marshal implements Meta.
func (s *std) Marshall(cl *Clue) ([]byte, error) {
	data := make(map[string]interface{})
	if cl.Meta != nil {
		second, err := json.Marshal(cl.Meta)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(second, &data)
		if err != nil {
			return nil, err
		}
	}
	data["data"] = nil
	if cl.Data != nil {
		var fieldData interface{}
		d, err := json.Marshal(cl.Data)
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
	s.Message = v
}

func MewStd(code, message string) Meta {
	return &std{
		Code:    code,
		Message: message,
	}
}
