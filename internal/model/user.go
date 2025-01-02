package model

import (
	"context"
	"net/mail"

	"github.com/pocketbase/pocketbase/core"
)

type User struct {
	ID        string
	Name      string
	Email     mail.Address
	Avatar    *string
}

func UserFromPBRecord(r *core.Record) *User {
	email, _ := mail.ParseAddress(r.GetString("email"))
	var avatar *string
  if a := r.GetString("avatar"); a != "" {
		avatar = &a
	}

	return &User{
		ID:     r.Id,
		Name:   r.GetString("name"),
		Email:  *email,
		Avatar: avatar,
	}
}

func UserFromContext(ctx context.Context) *User {
  u, _ := ctx.Value("user").(*User)
  return u
}
