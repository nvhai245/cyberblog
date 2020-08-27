package controller

import (
	"context"
	"github.com/nvhai245/cyberblog/server/cyber/connection"
	pb "github.com/nvhai245/cyberblog/server/cyber/proto"
	readPb "github.com/nvhai245/cyberblog/server/read/proto"
	writePb "github.com/nvhai245/cyberblog/server/write/proto"
	"log"
	"time"
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
		user := readUserToCyberUser(foundUser)
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

func EditUser(requestorEmail string, requestorIsAdmin bool, user *pb.User) *pb.EditUserResponse {
	log.Printf(`controller.EditUser(): [Request]: 
       %+v
       %+v
       %+v`, requestorEmail, requestorIsAdmin, user)

	user.UpdatedAt = time.Now().Unix()
	res, err := connection.WriteClient.EditUser(context.Background(), &writePb.EditUserRequest{
		RequestorEmail:   requestorEmail,
		RequestorIsAdmin: requestorIsAdmin,
		User:             cyberUserToWriteUser(user),
	})
	if err != nil {
		log.Println("Err in controller.EditUser(): ", err)
		return &pb.EditUserResponse{
			Success: false,
			User:    nil,
		}
	}
	if !res.GetSuccess() {
		return &pb.EditUserResponse{
			Success: false,
			User:    nil,
		}
	}
	response := &pb.EditUserResponse{
		Success: true,
		User:    writeUserToCyberUser(res.GetUser()),
	}
	// ***************************************************************************************************************
	log.Printf("controller.EditUser(): [Response]: %+v\n", response)
	return response
}

func readUserToCyberUser(foundUser *readPb.User) *pb.User {
	cyberUser := &pb.User{
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
	return cyberUser
}

func writeUserToCyberUser(foundUser *writePb.NewUser) *pb.User {
	cyberUser := &pb.User{
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
	return cyberUser
}

func cyberUserToWriteUser(userToWrite *pb.User) *writePb.NewUser {
	user := &writePb.NewUser{
		Id:        userToWrite.GetId(),
		Username:  userToWrite.GetUsername(),
		Email:     userToWrite.GetEmail(),
		FirstName: userToWrite.GetFirstName(),
		LastName:  userToWrite.GetLastName(),
		Avatar:    userToWrite.GetAvatar(),
		Birthday:  userToWrite.GetBirthday(),
		Bio:       userToWrite.GetBio(),
		Facebook:  userToWrite.GetFacebook(),
		Instagram: userToWrite.GetInstagram(),
		Twitter:   userToWrite.GetTwitter(),
		IsAdmin:   userToWrite.GetIsAdmin(),
		CreatedAt: userToWrite.GetCreatedAt(),
		UpdatedAt: userToWrite.GetUpdatedAt(),
	}
	return user
}
