package controller

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/nvhai245/cyberblog/server/auth/connection"
	pb "github.com/nvhai245/cyberblog/server/auth/proto"
	writePb "github.com/nvhai245/cyberblog/server/write/proto"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// Register controller fuc
func Register(req *pb.RegisterRequest) *pb.RegisterResponse {
	hashedPassword := hashAndSalt([]byte(req.GetPassword()))
	res, err := connection.WriteClient.SaveUser(context.Background(), &writePb.SaveUserRequest{
		User: &writePb.NewUser{
			Username:  "",
			Email:     req.GetEmail(),
			FirstName: "",
			LastName:  "",
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

	generatedJWT := generateJWT(savedUser.GetEmail())

	return &pb.RegisterResponse{
		User: &pb.SavedUser{
			Username:  savedUser.GetUsername(),
			Email:     savedUser.GetEmail(),
			FirstName: savedUser.GetFirstName(),
			LastName:  savedUser.GetLastName(),
			Avatar:    savedUser.GetAvatar(),
			Bio:       savedUser.GetBio(),
			Facebook:  savedUser.GetFacebook(),
			Instagram: savedUser.GetInstagram(),
			Twitter:   savedUser.GetTwitter(),
		},
		Token: generatedJWT,
	}
}

func Login(req *pb.LoginRequest) *pb.LoginResponse {
	res, err := connection.WriteClient.GetUser(context.Background(), &writePb.GetUserRequest{Email: req.GetEmail()})
	if err != nil {
		log.Printf("Error in controller.Login(), rpc call WriteClient.GetUser(): %v", err)
		return &pb.LoginResponse{}
	}

	if res.GetUser() == nil {
		log.Printf("Error in controller.Login(), Failed WriteClient.GetUser(): Can't find user with given email!")
		return &pb.LoginResponse{}
	}

	loggedInUser := res.GetUser()

	generatedJWT := generateJWT(loggedInUser.GetEmail())

	return &pb.LoginResponse{
		User: &pb.SavedUser{
			Username:  loggedInUser.GetUsername(),
			Email:     loggedInUser.GetEmail(),
			FirstName: loggedInUser.GetFirstName(),
			LastName:  loggedInUser.GetLastName(),
			Avatar:    loggedInUser.GetAvatar(),
			Bio:       loggedInUser.GetBio(),
			Facebook:  loggedInUser.GetFacebook(),
			Instagram: loggedInUser.GetInstagram(),
			Twitter:   loggedInUser.GetTwitter(),
		},
		Token: generatedJWT,
	}
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func generateJWT(email string) (jwtString string) {
	mySigningKey := []byte(os.Getenv("JWT_SECRET"))

	type MyCustomClaims struct {
		Email string `json:"email"`
		jwt.StandardClaims
	}

	claims := MyCustomClaims{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "auth",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Println("Error in controller.generateJWT(): ", err)
		return ""
	}
	return ss
}
