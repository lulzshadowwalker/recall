package template

import (
	"context"

	"github.com/lulzshadowwalker/recall/internal/model"
)

func user(ctx context.Context) *model.User {
  u := model.UserFromContext(ctx)
  return u
}
