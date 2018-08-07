package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/jotacamou/cngine/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
