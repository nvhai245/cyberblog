package main

import (
	"context"
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
	return res, nil
}