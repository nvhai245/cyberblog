package helper

import (
	"context"
	"github.com/gorilla/sessions"
	"net/http"
)

// HTTPKey is the key used to extract the Http struct.
type HTTPKey string

// HTTP is the struct used to inject the response writer and request http structs.
type HTTP struct {
	W *http.ResponseWriter
	R *http.Request
}

// GetSession returns a cached session of the given name
func GetSession(ctx context.Context, name string) *sessions.Session {
	store := ctx.Value(HTTPKey("session")).(*sessions.CookieStore)
	httpContext := ctx.Value(HTTPKey("http")).(HTTP)

	// Ignore err because a session is always returned even if one doesn't exist
	session, _ := store.Get(httpContext.R, name)

	return session
}

// SaveSession saves the session by writing it to the response
func SaveSession(ctx context.Context, session *sessions.Session) error {
	httpContext := ctx.Value(HTTPKey("http")).(HTTP)

	err := session.Save(httpContext.R, *httpContext.W)

	return err
}
