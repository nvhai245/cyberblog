package controller

import (
	"context"
	"github.com/nvhai245/cyberblog/server/cyber/connection"
	pb "github.com/nvhai245/cyberblog/server/cyber/proto"
	readPb "github.com/nvhai245/cyberblog/server/read/proto"
	"log"
)

func GetUserById(requestorId int32, userId int32) *pb.GetUserByIdResponse {
	log.Printf("controller.GetUserById(): [Request]: { userID: %v}", userId)
	// Check admin later
	_ = requestorId

	// ***************************************************************************************************************
	res, err := connection.ReadClient.GetUserById(context.Background(), &readPb.GetUserByIdRequest{UserId: userId})
	if err != nil {
		log.Println("Err in controller.GetUserById(): ", err)
		return &pb.GetUserByIdResponse{
			Success: false,
			User:    nil,
		}
	}
	foundUser := res.GetUser()
	if !res.GetSuccess() {
		return &pb.GetUserByIdResponse{
			Success: false,
			User:    nil,
		}
	}
	response := &pb.GetUserByIdResponse{
		Success: true,
		User: &pb.User{
			Id:        foundUser.GetId(),
			Username:  foundUser.GetUsername(),
			Email:     foundUser.GetEmail(),
			FirstName: foundUser.GetFirstName(),
			LastName:  foundUser.GetLastName(),
			Avatar:    foundUser.GetAvatar(),
			Birthday:  foundUser.GetBirthday(),
			Bio:       foundUser.GetBio(),
			Facebook:  foundUser.GetFacebook(),
			Instagram: foundUser.GetInstagram(),
			Twitter:   foundUser.GetTwitter(),
			IsAdmin:   foundUser.GetIsAdmin(),
			CreatedAt: foundUser.GetCreatedAt(),
			UpdatedAt: foundUser.GetUpdatedAt(),
		},
	}
	// ***************************************************************************************************************

	log.Printf("controller.GetUserById(): [Response]: %+v\n", response)
	return response
}

func GetAllUsers(adminId int32) *pb.GetAllUsersResponse {
	log.Printf("controller.GetUserById(): [Request]: { adminID: %v}", adminId)
	// ***************************************************************************************************************
	res, err := connection.ReadClient.GetAllUsers(context.Background(), &readPb.GetAllUsersRequest{AdminId: adminId})
	if err != nil {
		log.Println("Err in controller.GetUserById(): ", err)
		return &pb.GetAllUsersResponse{
			Success: false,
			Users:   nil,
		}
	}
	foundUsers := res.GetUsers()
	if !res.GetSuccess() {
		return &pb.GetAllUsersResponse{
			Success: false,
			Users:   nil,
		}
	}
	var users []*pb.User
	for _, foundUser := range foundUsers {
		user := &pb.User{
			Id:        foundUser.GetId(),
			Username:  foundUser.GetUsername(),
			Email:     foundUser.GetEmail(),
			FirstName: foundUser.GetFirstName(),
			LastName:  foundUser.GetLastName(),
			Avatar:    foundUser.GetAvatar(),
			Birthday:  foundUser.GetBirthday(),
			Bio:       foundUser.GetBio(),
			Facebook:  foundUser.GetFacebook(),
			Instagram: foundUser.GetInstagram(),
			Twitter:   foundUser.GetTwitter(),
			IsAdmin:   foundUser.GetIsAdmin(),
			CreatedAt: foundUser.GetCreatedAt(),
			UpdatedAt: foundUser.GetUpdatedAt(),
		}
		users = append(users, user)
	}
	response := &pb.GetAllUsersResponse{
		Success: true,
		Users:   users,
	}
	// ***************************************************************************************************************

	log.Printf("controller.GetUserById(): [Response]: %+v\n", response)
	return response
}