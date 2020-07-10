package main

import (
	"context"

	"github.com/nvhai245/cyberblog/server/auth/controller"
	pb "github.com/nvhai245/cyberblog/server/auth/proto"
)

// AuthServer for rpc
type AuthServer struct {
	pb.UnimplementedAuthServer
}

// Register method
func (*AuthServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	registeredUser := controller.Register(req)
	return registeredUser, nil
}

func (*AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	loggedInUser := controller.Login(req)
	return loggedInUser, nil
}
