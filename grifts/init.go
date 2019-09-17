package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/thejayhaykid/vuer/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
