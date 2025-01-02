package handler

import "net/http"

// hxRedirect sends a redirect response based on the HX-Request header.
//
// If the request is from HTMX (HX-Request header is present), it sets the HX-Redirect
// header and responds with a http see other status status.
//
// Otherwise, it performs a standard HTTP redirect.
func hxRedirect(w http.ResponseWriter, r *http.Request, to string) {
	if r.Header.Get("HX-Request") != "" {
		w.Header().Set("HX-Redirect", to)
		w.WriteHeader(http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, to, http.StatusSeeOther)
}
