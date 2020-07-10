package controller

import (
	"log"

	"github.com/nvhai245/cyberblog/server/write/model"
	pb "github.com/nvhai245/cyberblog/server/write/proto"
)

// SaveUser controller
func SaveUser(req *pb.SaveUserRequest) *pb.SaveUserResponse {
	newUser := req.GetUser()
	userToSave := &model.NewUser{
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
	}
	success := model.Insert(userToSave)

	if !success {
		log.Println("model.Insert() failed in controller.SaveUser()")
		return &pb.SaveUserResponse{
			User:    &pb.NewUser{},
			Success: false,
		}
	}

	return &pb.SaveUserResponse{
		User:    newUser,
		Success: true,
	}
}
