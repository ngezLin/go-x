package clue

import (
	"github.com/labstack/echo/v4"
)

type Sender interface {
	Send(c echo.Context) error
}

type sender struct {
	clue *Clue
}

func (s *sender) Send(c echo.Context) error {
	ctx := c.Request().Context()
	s.clue = s.clue.Meta.Templating(ctx, s.clue)
	return c.JSON(s.clue.HttpCode, s.clue)
}
