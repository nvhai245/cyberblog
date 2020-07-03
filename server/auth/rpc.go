package main

import (
	"context"

	pb "github.com/nvhai245/cyberblog/server/auth/proto"
	"github.com/nvhai245/cyberblog/server/auth/controller"

	"github.com/davecgh/go-spew/spew"
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