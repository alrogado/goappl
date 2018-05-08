package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/goappl/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
