package auth

import (
	"log/slog"

	"github.com/lulzshadowwalker/woosh/internal/http/cookie"
	"github.com/lulzshadowwalker/woosh/internal/model"
	"github.com/pocketbase/pocketbase/core"
)

func User(e *core.RequestEvent) *model.User {
	tok, ok := cookie.GetAuthCookie(e.Request)
	if !ok {
		return nil
	}

	rec, err := e.App.FindAuthRecordByToken(tok)
	if err != nil {
		slog.Error("failed to find auth record by token", "err", err)
		return nil
	}

	return model.UserFromPBRecord(rec)
}
