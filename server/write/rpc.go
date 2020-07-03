package main

import (
	"context"

	pb "github.com/nvhai245/cyberblog/server/write/proto"

	"github.com/davecgh/go-spew/spew"
)

// WriteServer for rpc 
type WriteServer struct {
	pb.UnimplementedWriteServer
}


// SaveUser method
func (*WriteServer) SaveUser(ctx context.Context, req *pb.SaveUserRequest) (*pb.SaveUserResponse, error) {
	
	return 
}