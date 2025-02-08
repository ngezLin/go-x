package clue

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type snap struct {
	Code    string `json:"responseCode"`
	Message string `json:"responseMessage"`
	*Clue
}

// GetCode implements Meta.
func (s *snap) GetCode() string {
	return s.Code
}

// GetInfo implements Meta.
func (s *snap) GetInfo() interface{} {
	return nil
}

// GetMessage implements Meta.
func (s *snap) GetMessage() string {
	return s.Message
}

// Templating implements Meta.
func (s *snap) Templating(ctx context.Context, clue *Clue) *Clue {
	//modified code
	var (
		code    int    = clue.HttpCode
		casee   string = clue.Meta.GetCode()
		service string = GetCtxServiceCode(ctx)
		message string = clue.Meta.GetMessage()
	)

	if casee == "" {
		casee = "00"
	}

	if service == "" {
		service = "00"
	}

	if message == "" && code == http.StatusOK {
		message = "Successful"
	}

	clue.Meta.SetCode(fmt.Sprintf("%d%s%s", code, service, casee))
	clue.Meta.SetMessage(message)

	return clue
}

// Marshal implements Meta.
func (s *snap) MarshalJSON() ([]byte, error) {
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
	if s.Clue.Data != nil {
		second, err := json.Marshal(s.Clue.Data)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(second, &data)
		if err != nil {
			return nil, err
		}
	}
	return json.Marshal(data)
}

// SetCode implements Meta.
func (s *snap) SetCode(v string) {
	s.Code = v
}

// SetMessage implements Meta.
func (s *snap) SetMessage(v string) {
	s.Code = v
}

func MewSnapBI(code, message string) Meta {
	return &snap{
		Code:    code,
		Message: message,
	}
}

const serviceCode = "snap-service-code"

func DefineCtxServiceCode(ctx context.Context, code string) context.Context {
	ctx = context.WithValue(ctx, serviceCode, code)
	return ctx
}

func GetCtxServiceCode(ctx context.Context) string {
	v, ok := ctx.Value(serviceCode).(string)
	if !ok {
		return ""
	}
	return v
}
