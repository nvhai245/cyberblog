package controller

import (
	"log"

	"github.com/nvhai245/cyberblog/server/write/model"
	pb "github.com/nvhai245/cyberblog/server/write/proto"
)

// SaveUser controller
func SaveUser(req *pb.SaveUserRequest) *pb.SaveUserResponse {
	newUser := req.GetUser()
	userToSave := &model.User{
		Username:  newUser.GetUsername(),
		Email:     newUser.GetEmail(),
		Hash:      req.GetHash(),
		FirstName: newUser.GetFirstName(),
		LastName:  newUser.GetLastName(),
		Avatar:    newUser.GetAvatar(),
		Birthday:  newUser.GetBirthday(),
		Bio:       newUser.GetBio(),
		Facebook:  newUser.GetFacebook(),
		Instagram: newUser.GetInstagram(),
		Twitter:   newUser.GetTwitter(),
		CreatedAt: newUser.GetCreatedAt(),
	}
	success, userID := model.Insert(userToSave)

	if !success {
		log.Println("controller.SaveUser() failed in model.Insert()")
		return &pb.SaveUserResponse{
			User:    &pb.NewUser{},
			Success: false,
		}
	}

	newUser.Id = userID

	return &pb.SaveUserResponse{
		User:    newUser,
		Success: true,
	}
}
