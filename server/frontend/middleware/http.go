package middleware

import (
	"context"
	"github.com/gorilla/sessions"
	"net/http"

	"github.com/nvhai245/cyberblog/server/frontend/helper"
)

// InjectHTTPMiddleware handles injecting the ResponseWriter and Request structs
// into context so that resolver methods can use these to set and read cookies. It also passes a // CookieStore initialized in `server.go` into context for facilitated cookie handling.
func InjectHTTPMiddleware(session *sessions.CookieStore) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			httpContext := helper.HTTP{
				W: &w,
				R: r,
			}
			httpKeyContext := helper.HTTPKey("http")

			sessionKeyContext := helper.HTTPKey("session")

			ctx := context.WithValue(r.Context(), httpKeyContext, httpContext)
			ctx = context.WithValue(ctx, sessionKeyContext, session)
			r = r.WithContext(ctx)

			// CORS
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
