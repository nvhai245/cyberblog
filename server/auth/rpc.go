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

// Register rpc
func (*AuthServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	registeredUser := controller.Register(req)
	return registeredUser, nil
}

// Login rpc
func (*AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	loggedInUser := controller.Login(req)
	return loggedInUser, nil
}

// CheckToken rpc
func (*AuthServer) CheckToken(ctx context.Context, req *pb.CheckTokenRequest) (*pb.CheckTokenResponse, error) {
	valid := controller.CheckToken(req)
	return valid, nil
}
