package main

import (
	"context"
	"github.com/nvhai245/cyberblog/server/read/controller"
	pb "github.com/nvhai245/cyberblog/server/read/proto"
	"log"
	_ "log"
)

// WriteServer for rpc
type ReadServer struct {
	pb.UnimplementedReadServer
}

func (*ReadServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	res := controller.GetUser(req)
	if res.GetUser() == nil {
		log.Println("controller.GetUser(): Can't find user with given email!")
	}
	return res, nil
}

func (*ReadServer) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
	res := controller.GetUserById(req)
	if res.GetSuccess() == false {
		log.Println("controller.GetUserById(): Can't find user with given id!")
	}
	return res, nil
}
