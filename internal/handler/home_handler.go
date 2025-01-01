package handler

import (
	"github.com/lulzshadowwalker/woosh/internal/template"
	"github.com/pocketbase/pocketbase/core"
)

func HandleGetHome(e *core.RequestEvent) error {
	return render(e, template.HomeIndex())
}
