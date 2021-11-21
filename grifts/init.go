package grifts

import (
	"CloudTechnology/NGHCI/sangformall/actions"

	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
