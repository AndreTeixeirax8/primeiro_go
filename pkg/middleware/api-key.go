package middleware

import (
	"net/http"
	"os"

	"github.com/go-chi/render"
)

func ApiKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")
		if apiKey != os.Getenv("API_KEY") {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{"error": "Unauthorized"})
			return
		}
		next.ServeHTTP(w, r.WithContext(r.Context()))
	})
}
