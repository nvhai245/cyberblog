package controller

import (
	"log"

	userModel "github.com/nvhai245/cyberblog/server/write/model/user"
	pb "github.com/nvhai245/cyberblog/server/write/proto"
)

// SaveUser controller
func SaveUser(req *pb.SaveUserRequest) *pb.SaveUserResponse {
	newUser := req.GetUser()
	userToSave := &userModel.User{
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
		IsAdmin:   newUser.GetIsAdmin(),
		CreatedAt: newUser.GetCreatedAt(),
	}
	success, userID := userModel.Insert(userToSave)

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

// EditUser controller
func EditUser(req *pb.EditUserRequest) *pb.EditUserResponse {
	userToEdit := protoUserToModelUser(req.GetUser())
	if !req.GetRequestorIsAdmin() {
		userToEdit.Email = req.GetRequestorEmail()
	}
	success, editedUser := userModel.Update(req.GetRequestorEmail(), req.GetRequestorIsAdmin(), userToEdit)
	if !success {
		log.Println("controller.Edit() failed in userModel.Update()")
		return &pb.EditUserResponse{
			Success: false,
			User:    &pb.NewUser{},
		}
	}
	userToReturn := modelUserToProtoUser(editedUser)
	return &pb.EditUserResponse{
		Success: true,
		User:    userToReturn,
	}
}

// protoUserToModelUser func
func protoUserToModelUser(newUser *pb.NewUser) *userModel.User {
	user := &userModel.User{
		ID:       newUser.GetId(),
		Username: newUser.GetUsername(),
		Email:    newUser.GetEmail(),
		//Hash:      req.GetHash(),
		FirstName: newUser.GetFirstName(),
		LastName:  newUser.GetLastName(),
		Avatar:    newUser.GetAvatar(),
		Birthday:  newUser.GetBirthday(),
		Bio:       newUser.GetBio(),
		Facebook:  newUser.GetFacebook(),
		Instagram: newUser.GetInstagram(),
		Twitter:   newUser.GetTwitter(),
		CreatedAt: newUser.GetCreatedAt(),
		UpdatedAt: newUser.GetUpdatedAt(),
	}
	return user
}

// modelUserToProtoUser func
func modelUserToProtoUser(foundUser *userModel.User) *pb.NewUser {
	user := &pb.NewUser{
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
	}
	return user
}
