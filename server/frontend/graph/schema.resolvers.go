package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/nvhai245/cyberblog/server/frontend/graph/generated"
	"github.com/nvhai245/cyberblog/server/frontend/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	// Call gRPC to Auth
	return &model.User{
		FirstName: "",
		LastName: "",
		Email: input.Email,
		CreatedAt: int(time.Now().Unix()),
	}, nil
}

func (r *mutationResolver) Register(ctx context.Context, email string, password string) (*model.Token, error) {
	expiredAt := time.Now().Add(time.Hour * 1).Unix()
	token := &model.Token{
		Token: "jwtToken@abcdef",
		ExpireAt: int(expiredAt),
	}
	return token, nil
}

func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*model.Token, error) {
	// Call gRPC to Auth
	expiredAt := time.Now().Add(time.Hour * 1).Unix()
	token := &model.Token{
		Token: "jwtToken@abcdef",
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
