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
