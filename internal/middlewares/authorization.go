package middlewares

import "net/http"

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Authorization logic
		next.ServeHTTP(w, r)
	})
}