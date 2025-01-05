package cookie

import (
	"errors"
	"log/slog"
	"net/http"
)

const authKey = "recall-auth-token"

var (
	ErrEmptyToken = errors.New("access token cannot be empty")
)

func SetAuthCookie(w http.ResponseWriter, token string) error {
	if token == "" {
		return ErrEmptyToken
	}

	http.SetCookie(w, &http.Cookie{
		Name:     authKey,
		Value:    token,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
		HttpOnly: true,
		Path:     "/",
	})

	return nil
}

func GetAuthCookie(r *http.Request) (token string, ok bool) {
	cookie, err := r.Cookie(authKey)
	if err != nil {
		if !errors.Is(err, http.ErrNoCookie) {
			slog.Error("failed to get auth cookie", "err", err)
		}

		return "", false
	}

	return cookie.Value, true
}

func ClearAuthCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     authKey,
		Value:    "",
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
		HttpOnly: true,
		Path:     "/",
		MaxAge:   -1,
	})
}
