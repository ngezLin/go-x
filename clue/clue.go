package clue

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
