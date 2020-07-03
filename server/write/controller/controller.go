package controller

import (
	"log"

	"github.com/nvhai245/cyberblog/server/read/connection"
	pb "github.com/nvhai245/cyberblog/server/write/proto"

	"google.golang.org/grpc"
)

// SaveUser controller
func SaveUser(req *pb.SaveUserRequest) (*pb.SaveUserResponse) {
	newUser := req.GetUser()
	
}