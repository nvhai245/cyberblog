package main

import (
	"context"
	"github.com/nvhai245/cyberblog/server/cyber/controller"
	pb "github.com/nvhai245/cyberblog/server/cyber/proto"
)

type CyberServer struct {
	pb.UnimplementedCyberServer
}

func (*CyberServer) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
	foundUser := controller.GetUserById(req.GetRequestorId(), req.GetUserId())
	return foundUser, nil
}

func (*CyberServer) GetAllUsers(ctx context.Context, req *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
	foundUser := controller.GetAllUsers(req.GetAdminId())
	return foundUser, nil
}

func (*CyberServer) EditUser(ctx context.Context, req *pb.EditUserRequest) (*pb.EditUserResponse, error) {
	editedUser := controller.EditUser(req.GetRequestorEmail(), req.GetRequestorIsAdmin(), req.GetUser())
	return editedUser, nil
}

func (*CyberServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	deletedUser, err := controller.DeleteUser(req.GetRequestorEmail(), req.GetRequestorIsAdmin(), req.GetUserId())
	return deletedUser, err
}
