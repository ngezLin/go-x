package clue

import "net/http"

type Clue struct {
	HttpCode int         `json:"-"`
	Data     interface{} `json:"-"`
	Meta     Meta        `json:"-"`
}

type Builder interface {
	Std() Builder
	SnapBI() Builder
	Sender
}

type builder struct {
	sender
}

func (b *builder) Error() string {
	return b.clue.Meta.GetMessage()
}

// SnapBI implements Builder.
func (b *builder) SnapBI() Builder {
	b.clue.Meta = MewSnapBI(b.clue.Meta.GetCode(), b.clue.Meta.GetMessage())
	return b
}

// Std implements Builder.
func (b *builder) Std() Builder {
	b.clue.Meta = MewStd(b.clue.Meta.GetCode(), b.clue.Meta.GetMessage())
	return b
}

func (b *builder) MarshalJSON() ([]byte, error) {
	return b.clue.Meta.Marshall(b.clue)
}

func Build(httpCode int, code string, data interface{}, message string) Builder {
	return &builder{
		sender: sender{
			clue: &Clue{
				HttpCode: httpCode,
				Meta: &std{
					Code:    code,
					Message: message,
				},
				Data: data,
			},
		},
	}
}

func CoverBuilder(err error, data interface{}) Builder {
	re, ok := err.(*builder)
	if ok {
		re.clue.Data = data
		return re
	} else {
		return Build(http.StatusInternalServerError, "00", data, err.Error())
	}
}
