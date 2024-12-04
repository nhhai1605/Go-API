package middlewares

import (
	"go-api/core"
	"go-api/internal/databases/auth"
	"net/http"
	"slices"

	"github.com/go-chi/render"
)

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//get "token" from the request header
		token := r.Header.Get("token")
		listToken := []string{}
		for _, v := range auth.MockTokens {
			listToken = append(listToken, v)
		}
		if !slices.Contains(listToken, token) {
			render.Render(w, r, core.ErrUnauthorized(nil))
			return
		}
		next.ServeHTTP(w, r)
	})
}