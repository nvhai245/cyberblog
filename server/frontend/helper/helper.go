package helper

import (
	"context"
	"github.com/gorilla/sessions"
	authPb "github.com/nvhai245/cyberblog/server/auth/proto"
	"github.com/nvhai245/cyberblog/server/frontend/graph/model"
	readPb "github.com/nvhai245/cyberblog/server/read/proto"
	writePb "github.com/nvhai245/cyberblog/server/write/proto"
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
	session.Options.HttpOnly = true

	err := session.Save(httpContext.R, *httpContext.W)

	return err
}

// ReadUserToGraphUser parse user object from cyber into graphql user
func ReadUserToGraphUser(foundUser *readPb.User) *model.User {
	user := &model.User{
		ID:        int(foundUser.GetId()),
		Username:  foundUser.GetUsername(),
		Email:     foundUser.GetEmail(),
		FirstName: foundUser.GetFirstName(),
		LastName:  foundUser.GetLastName(),
		Avatar:    foundUser.GetAvatar(),
		Birthday:  int(foundUser.GetBirthday()),
		Bio:       foundUser.GetBio(),
		Facebook:  foundUser.GetFacebook(),
		Instagram: foundUser.GetInstagram(),
		Twitter:   foundUser.GetTwitter(),
		IsAdmin:   foundUser.GetIsAdmin(),
		CreatedAt: int(foundUser.GetCreatedAt()),
		UpdatedAt: int(foundUser.GetUpdatedAt()),
	}
	return user
}

// WriteUserToGraphUser parse user object from cyber into graphql user
func WriteUserToGraphUser(foundUser *writePb.NewUser) *model.User {
	user := &model.User{
		ID:        int(foundUser.GetId()),
		Username:  foundUser.GetUsername(),
		Email:     foundUser.GetEmail(),
		FirstName: foundUser.GetFirstName(),
		LastName:  foundUser.GetLastName(),
		Avatar:    foundUser.GetAvatar(),
		Birthday:  int(foundUser.GetBirthday()),
		Bio:       foundUser.GetBio(),
		Facebook:  foundUser.GetFacebook(),
		Instagram: foundUser.GetInstagram(),
		Twitter:   foundUser.GetTwitter(),
		IsAdmin:   foundUser.GetIsAdmin(),
		CreatedAt: int(foundUser.GetCreatedAt()),
		UpdatedAt: int(foundUser.GetUpdatedAt()),
	}
	return user
}

// AuthUserToGraphUser parse user object from cyber into graphql user
func AuthUserToGraphUser(foundUser *authPb.SavedUser) *model.User {
	user := &model.User{
		ID:        int(foundUser.GetId()),
		Username:  foundUser.GetUsername(),
		Email:     foundUser.GetEmail(),
		FirstName: foundUser.GetFirstName(),
		LastName:  foundUser.GetLastName(),
		Avatar:    foundUser.GetAvatar(),
		Birthday:  int(foundUser.GetBirthday()),
		Bio:       foundUser.GetBio(),
		Facebook:  foundUser.GetFacebook(),
		Instagram: foundUser.GetInstagram(),
		Twitter:   foundUser.GetTwitter(),
		IsAdmin:   foundUser.GetIsAdmin(),
		CreatedAt: int(foundUser.GetCreatedAt()),
		UpdatedAt: int(foundUser.GetUpdatedAt()),
	}
	return user
}
