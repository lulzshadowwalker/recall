package handler

import (
	"github.com/a-h/templ"
	"github.com/pocketbase/pocketbase/core"
)

func render(e *core.RequestEvent, c templ.Component) error {
	return c.Render(e.Request.Context(), e.Response)
}
