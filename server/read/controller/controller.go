package controller

import (
	"github.com/nvhai245/cyberblog/server/read/model"
	pb "github.com/nvhai245/cyberblog/server/read/proto"
	"log"
)

// GetUser controller
func GetUser(req *pb.GetUserRequest) *pb.GetUserResponse {
	foundUser := model.GetUserByEmail(req.GetEmail())
	if foundUser == nil {
		log.Println("controller.GetUser() failed in model.GetUserByEmail()")
	}
	return &pb.GetUserResponse{
		User: &pb.User{
			Id:        foundUser.ID,
			Username:  foundUser.Username,
			Email:     foundUser.Email,
			FirstName: foundUser.FirstName,
			LastName:  foundUser.LastName,
			Avatar:    foundUser.Avatar,
			Birthday:  foundUser.Birthday,
			Bio:       foundUser.Bio,
			Facebook:  foundUser.Facebook,
			Instagram: foundUser.Instagram,
			Twitter:   foundUser.Twitter,
			IsAdmin:   foundUser.IsAdmin,
			CreatedAt: foundUser.CreatedAt,
			UpdatedAt: foundUser.UpdatedAt,
		},
		Hash: foundUser.Hash,
	}
}

// GetUserById controller
func GetUserById(req *pb.GetUserByIdRequest) *pb.GetUserByIdResponse {
	foundUser := model.GetUserById(req.GetUserId())
	if foundUser == nil {
		log.Println("controller.GetUserById() failed in model.GetUserById()")
		return &pb.GetUserByIdResponse{
			Success: false,
			User:    nil,
		}
	}
	return &pb.GetUserByIdResponse{
		User: &pb.User{
			Id:        foundUser.ID,
			Username:  foundUser.Username,
			Email:     foundUser.Email,
			FirstName: foundUser.FirstName,
			LastName:  foundUser.LastName,
			Avatar:    foundUser.Avatar,
			Birthday:  foundUser.Birthday,
			Bio:       foundUser.Bio,
			Facebook:  foundUser.Facebook,
			Instagram: foundUser.Instagram,
			Twitter:   foundUser.Twitter,
			IsAdmin:   foundUser.IsAdmin,
			CreatedAt: foundUser.CreatedAt,
			UpdatedAt: foundUser.UpdatedAt,
		},
		Success: true,
	}
}
