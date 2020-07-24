package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	cyberPb "github.com/nvhai245/cyberblog/server/cyber/proto"
	"log"

	authPb "github.com/nvhai245/cyberblog/server/auth/proto"
	"github.com/nvhai245/cyberblog/server/frontend/connection"
	"github.com/nvhai245/cyberblog/server/frontend/graph/generated"
	"github.com/nvhai245/cyberblog/server/frontend/graph/model"
	"github.com/nvhai245/cyberblog/server/frontend/helper"
)

func (r *mutationResolver) Register(ctx context.Context, email string, password string) (*model.AuthResponse, error) {
	res, err := connection.AuthClient.Register(context.Background(), &authPb.RegisterRequest{Email: email, Password: password})
	if err != nil {
		log.Println("Error in rpc AuthClient.Register(): ", err)
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	//expiredAt := time.Now().Add(time.Hour * 24).Unix()
	session := helper.GetSession(ctx, "auth")
	session.Values["token"] = res.GetToken()
	if err := helper.SaveSession(ctx, session); err != nil {
		return nil, fmt.Errorf("Resolver.Register(): Failed to save cart in session with error: %s", err)
	}

	registeredUser := res.GetUser()
	var user *model.User
	var response *model.AuthResponse

	if registeredUser == nil {
		return nil, fmt.Errorf("AN USER WITH GIVEN CREDENTIALS ALREADY EXIST!")
	}

	user = &model.User{
		ID:        int(registeredUser.GetId()),
		Username:  registeredUser.GetUsername(),
		Email:     registeredUser.GetEmail(),
		FirstName: registeredUser.GetFirstName(),
		LastName:  registeredUser.GetLastName(),
		Avatar:    registeredUser.GetAvatar(),
		Birthday:  int(registeredUser.GetBirthday()),
		Bio:       registeredUser.GetBio(),
		Facebook:  registeredUser.GetFacebook(),
		Instagram: registeredUser.GetInstagram(),
		Twitter:   registeredUser.GetTwitter(),
		IsAdmin:   registeredUser.GetIsAdmin(),
		CreatedAt: int(registeredUser.GetCreatedAt()),
	}
	response = &model.AuthResponse{
		Message: "REGISTERED SUCCESSFUL!",
		User:    user,
	}
	// ****************************************************************************************************************

	log.Printf("Called AuthClient.Register() successful!, reply: %+v\n", response)
	return response, nil
}

func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*model.AuthResponse, error) {
	res, err := connection.AuthClient.Login(context.Background(), &authPb.LoginRequest{Email: email, Password: password})
	if err != nil {
		log.Println("Error in rpc AuthClient.Login(): ", err)
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	//expiredAt := time.Now().Add(time.Hour * 24).Unix()
	session := helper.GetSession(ctx, "auth")
	session.Values["token"] = res.GetToken()
	if err := helper.SaveSession(ctx, session); err != nil {
		return nil, fmt.Errorf("Resolver.Register(): Failed to save cart in session with error: %s", err)
	}

	loggedInUser := res.GetUser()
	var user *model.User

	if loggedInUser == nil {
		return nil, fmt.Errorf("WRONG EMAIL OR PASSWORD!")
	}
	user = &model.User{
		ID:        int(loggedInUser.GetId()),
		Username:  loggedInUser.GetUsername(),
		Email:     loggedInUser.GetEmail(),
		FirstName: loggedInUser.GetFirstName(),
		LastName:  loggedInUser.GetLastName(),
		Avatar:    loggedInUser.GetAvatar(),
		Birthday:  int(loggedInUser.GetBirthday()),
		Bio:       loggedInUser.GetBio(),
		Facebook:  loggedInUser.GetFacebook(),
		Instagram: loggedInUser.GetInstagram(),
		Twitter:   loggedInUser.GetTwitter(),
		IsAdmin:   loggedInUser.GetIsAdmin(),
		CreatedAt: int(loggedInUser.GetCreatedAt()),
		UpdatedAt: int(loggedInUser.GetUpdatedAt()),
	}
	response := &model.AuthResponse{
		Message: "LOGGED IN!",
		User:    user,
	}
	// ****************************************************************************************************************

	log.Printf("Called AuthClient.Login() successful!, reply: %+v\n", response)
	return response, nil
}

func (r *mutationResolver) GetUserByID(ctx context.Context, requestorID int, userID int) (*model.GetUserByIDResponse, error) {
	session := helper.GetSession(ctx, "auth")
	token := fmt.Sprintf("%v", session.Values["token"])
	checkResponse, err := connection.AuthClient.CheckToken(context.Background(), &authPb.CheckTokenRequest{Token: token})
	if err != nil {
		log.Println("Error in rpc AuthClient.CheckToken(): ", err)
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	if !checkResponse.GetValid() {
		return nil, fmt.Errorf("UNAUTHORIZED!")
	}
	res, err := connection.CyberClient.GetUserById(context.Background(), &cyberPb.GetUserByIdRequest{
		RequestorId: int32(requestorID),
		UserId:      int32(userID),
	})
	if err != nil {
		log.Println("Error in rpc CyberClient.GetUserById(): ", err)
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	foundUser := res.GetUser()
	if foundUser == nil {
		return nil, fmt.Errorf("Can't find user with given id!")
	}
	response := &model.GetUserByIDResponse{
		Message: "User found!",
		User: &model.User{
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
		},
	}
	// ****************************************************************************************************************

	log.Printf("Called CyberClient.GetUserById() successful!, reply: %+v\n", response)
	return response, nil
}

func (r *mutationResolver) GetAllUsers(ctx context.Context, adminID int) (*model.GetAllUsersResponse, error) {
	session := helper.GetSession(ctx, "auth")
	token := fmt.Sprintf("%v", session.Values["token"])
	checkResponse, err := connection.AuthClient.CheckToken(context.Background(), &authPb.CheckTokenRequest{Token: token})
	if err != nil {
		log.Println("Error in rpc AuthClient.CheckToken(): ", err)
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	if !checkResponse.GetValid() {
		return nil, fmt.Errorf("UNAUTHORIZED!")
	}
	if !checkResponse.GetIsAdmin() {
		return nil, fmt.Errorf("YOU ARE NOT AN ADMIN!")
	}

	res, err := connection.CyberClient.GetAllUsers(context.Background(), &cyberPb.GetAllUsersRequest{
		AdminId: int32(adminID),
	})
	if err != nil {
		log.Println("Error in rpc CyberClient.GetAllUsers(): ", err)
		return nil, fmt.Errorf("INTERNAL SERVER ERROR!")
	}
	foundUsers := res.GetUsers()
	if foundUsers == nil {
		return nil, fmt.Errorf("Can't get all users!")
	}
	var users []*model.User
	for _, foundUser := range foundUsers {
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
		users = append(users, user)
	}
	response := &model.GetAllUsersResponse{
		Message: "Get all users successful!",
		Users:   users,
	}

	// ****************************************************************************************************************

	log.Printf("Called CyberClient.GetAllUsers() successful!, reply: %+v\n", response)
	return response, nil
}

func (r *mutationResolver) EditUser(ctx context.Context, userID int, editedUser model.EditedUser) (*model.EditUserResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteUser(ctx context.Context, adminID int, userID int) (*model.DeleteUserResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return []*model.User{}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
