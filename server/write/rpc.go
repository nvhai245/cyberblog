package main

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"

	"github.com/nvhai245/cyberblog/server/write/controller"
	pb "github.com/nvhai245/cyberblog/server/write/proto"
)

// WriteServer for rpc
type WriteServer struct {
	pb.UnimplementedWriteServer
}

// SaveUser method
func (*WriteServer) SaveUser(ctx context.Context, req *pb.SaveUserRequest) (*pb.SaveUserResponse, error) {
	res := controller.SaveUser(req)
	if res.GetSuccess() == false {
		log.Println("controller.SaveUser(): success: false")
	}
	return res, nil
}

func (*WriteServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	res := controller.GetUser(req)
	if res.GetUser() == nil {
		log.Println("controller.GetUser(): Can't find user with given email!")
	}
	return res, nil
}
