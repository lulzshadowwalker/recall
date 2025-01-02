package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/lulzshadowwalker/woosh/internal/handler"
	"github.com/lulzshadowwalker/woosh/internal/http/middleware"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// serves static files from the provided public dir (if exists)
		se.Router.GET("/public/{path...}", apis.Static(os.DirFS("./pb_public"), false))

		se.Router.GET("/welcome", handler.HandleGetWelcome)
		se.Router.GET("/", handler.HandleGetHome).BindFunc(middleware.Authentication)

    se.Router.POST("/auth/logout", handler.HandleAuthLogout)

		return se.Next()
	})

	app.OnRecordAuthWithOAuth2Request().BindFunc(func(e *core.RecordAuthWithOAuth2RequestEvent) error {
		token, err := e.Auth.NewAuthToken()
		if err != nil {
			return fmt.Errorf("failed to create new auth token %w", err)
		}

		const authTokenKey = "woosh-auth-token"

		e.SetCookie(&http.Cookie{
			Name:     authTokenKey,
			Value:    token,
			SameSite: http.SameSiteStrictMode,
			Secure:   true,
			HttpOnly: true,
			Path:     "/",
		})

		return e.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
