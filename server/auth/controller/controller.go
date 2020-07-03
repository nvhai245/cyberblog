package controller

import (
	"context"
	"log"

	pb "github.com/nvhai245/cyberblog/server/auth/proto"
	writePb "github.com/nvhai245/cyberblog/server/write/proto"
	"github.com/nvhai245/cyberblog/server/auth/connection"

	"github.com/davecgh/go-spew/spew"
	"golang.org/x/crypto/bcrypt"
)

// Register controller fuc
func Register(req *pb.RegisterRequest) (*pb.RegisterResponse) {
	hashedPassword := hashAndSalt([]byte(req.GetPassword()))
	res, err := connection.WriteClient.SaveUser(context.Background(), &writePb.SaveUserRequest{
		User: &writePb.NewUser{
			Username: "",
			Email: req.GetEmail(),
			FirstName: "",
			LastName: "",
		},
		Hash: hashedPassword,
	})
	if err != nil {
		log.Printf("Error in controller.Register(), rpc call WriteClient.SaveUser(): %v", err)
		return &pb.RegisterResponse{}
	}

	if !res.GetSuccess() {
		log.Printf("Error in controller.Register(), Failed WriteClient.SaveUser()")
		return &pb.RegisterResponse{}
	}

	savedUser := res.GetUser()
	return &pb.RegisterResponse{
		User: &pb.SavedUser{
			Username: savedUser.GetUsername(),
			Email: savedUser.GetEmail(),
			FirstName: savedUser.GetFirstName(),
			LastName: savedUser.GetLastName(),
			Avatar: savedUser.GetAvatar(),
			Bio: savedUser.GetBio(),
			Facebook: savedUser.GetFacebook(),
			Instagram: savedUser.GetInstagram(),
			Twitter: savedUser.GetTwitter(),
		},
	}
}

func hashAndSalt(pwd []byte) string {
    hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
    if err != nil {
        log.Println(err)
    }
    return string(hash)
}