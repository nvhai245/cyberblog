package main

import (
	"context"
	"github.com/nvhai245/cyberblog/server/write/controller/user"
	pb "github.com/nvhai245/cyberblog/server/write/proto"
)

// WriteServer for rpc
type WriteServer struct {
	pb.UnimplementedWriteServer
}

// SaveUser method
func (*WriteServer) SaveUser(ctx context.Context, req *pb.SaveUserRequest) (*pb.SaveUserResponse, error) {
	res := userController.SaveUser(req)
	return res, nil
}

// EditUser method
func (*WriteServer) EditUser(ctx context.Context, req *pb.EditUserRequest) (*pb.EditUserResponse, error) {
	res := userController.EditUser(req)
	return res, nil
}

// DeleteUser method
func (*WriteServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	res, err := userController.DeleteUser(req)
	return res, err
}
