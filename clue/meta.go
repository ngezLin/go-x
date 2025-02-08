package clue

import "context"

type Meta interface {
	GetMessage() string
	SetMessage(string)
	SetCode(string)
	GetCode() string
	GetInfo() interface{}
	Templating(ctx context.Context, clue *Clue) *Clue
	Marshall(clue *Clue) ([]byte, error)
}
