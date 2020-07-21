package controller

import (
	pb "github.com/nvhai245/cyberblog/server/cyber/proto"
)

func GetUserById(requestorId int32, userId int32) *pb.GetUserByIdResponse {
	_ = requestorId
	_ = userId
	return &pb.GetUserByIdResponse{
		Success: false,
		User:    &pb.User{},
	}
}
