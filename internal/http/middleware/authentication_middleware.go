package middleware

import (
	"context"
	"net/http"

	"github.com/lulzshadowwalker/recall/internal/http/auth"
	"github.com/pocketbase/pocketbase/core"
)

func Authentication(e *core.RequestEvent) error {
	if auth.User(e) == nil {
		return e.Redirect(http.StatusSeeOther, "/welcome")
	}

  ctx := context.WithValue(e.Request.Context(), "user", auth.User(e))
  e.Request = e.Request.WithContext(ctx)

  return e.Next()
}
