package handler

import (
	"github.com/lulzshadowwalker/woosh/internal/template"
	"github.com/pocketbase/pocketbase/core"
)

func HandleGetWelcome(e *core.RequestEvent) error {
	return render(e, template.WelcomeIndex())
}
