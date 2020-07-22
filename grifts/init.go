package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/njncalub/coke/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
