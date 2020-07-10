package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/nvhai245/cyberblog/server/frontend/helper"
	"log"
	"time"

	authPb "github.com/nvhai245/cyberblog/server/auth/proto"
	"github.com/nvhai245/cyberblog/server/frontend/connection"
	"github.com/nvhai245/cyberblog/server/frontend/graph/generated"
	"github.com/nvhai245/cyberblog/server/frontend/graph/model"
)

func (r *mutationResolver) Register(ctx context.Context, email string, password string) (*model.Token, error) {
	res, err := connection.AuthClient.Register(context.Background(), &authPb.RegisterRequest{Email: email, Password: password})
	if err != nil {
		log.Println("Error in rpc AuthClient.Register: ", err)
		return &model.Token{}, nil
	}
	log.Println("Called AuthClient.Register() successful!")
	expiredAt := time.Now().Add(time.Hour * 24).Unix()

	// Grab a session (gorilla returns a session if the named session doesn't exist)
	session := helper.GetSession(ctx, "auth")

	// Reading userID cookie value
	token := session.Values["token"]
	log.Println("current token is: ", token)

	// Setting userID cookie value
	session.Values["token"] = res.GetToken()

	// Save session
	if err := helper.SaveSession(ctx, session); err != nil {
		return nil, fmt.Errorf("Failed to save cart in session with error: %s", err)
	}

	return &model.Token{
		Token:    res.GetToken(),
		ExpireAt: int(expiredAt),
	}, nil

}

func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*model.Token, error) {
	// Call gRPC to Auth
	expiredAt := time.Now().Add(time.Hour * 1).Unix()
	token := &model.Token{
		Token:    "jwtToken@abcdef",
		ExpireAt: int(expiredAt),
	}
	return token, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	// Call gRPC to Biz
	return []*model.User{}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
