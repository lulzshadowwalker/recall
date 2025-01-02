package handler

import (
	"github.com/lulzshadowwalker/woosh/internal/http/cookie"
	"github.com/pocketbase/pocketbase/core"
)

func HandleAuthLogout(e *core.RequestEvent) error {
  cookie.ClearAuthCookie(e.Response)
  hxRedirect(e.Response, e.Request, "/welcome")

  return nil
}
