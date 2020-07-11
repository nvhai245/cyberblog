package controller

import (
	"context"
	readPb "github.com/nvhai245/cyberblog/server/read/proto"
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
	log.Printf("controller.Register(): [Request]: %+v\n", req)

	// ***************************************************************************************************************

	hashedPassword := hashAndSalt([]byte(req.GetPassword()))
	res, err := connection.WriteClient.SaveUser(context.Background(), &writePb.SaveUserRequest{
		User: &writePb.NewUser{
			Username:  "",
			Email:     req.GetEmail(),
			FirstName: "",
			LastName:  "",
			CreatedAt: time.Now().Unix(),
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
	response := &pb.RegisterResponse{
		User: &pb.SavedUser{
			Id:        savedUser.GetId(),
			Username:  savedUser.GetUsername(),
			Email:     savedUser.GetEmail(),
			FirstName: savedUser.GetFirstName(),
			LastName:  savedUser.GetLastName(),
			Avatar:    savedUser.GetAvatar(),
			Bio:       savedUser.GetBio(),
			Facebook:  savedUser.GetFacebook(),
			Instagram: savedUser.GetInstagram(),
			Twitter:   savedUser.GetTwitter(),
			CreatedAt: savedUser.GetCreatedAt(),
		},
		Token: generatedJWT,
	}

	// ***************************************************************************************************************

	log.Printf("controller.Register(): [Response]: %+v\n", response)
	return response
}

func Login(req *pb.LoginRequest) *pb.LoginResponse {
	log.Printf("controller.Login(): [Request]: %+v\n", req)

	// ***************************************************************************************************************

	res, err := connection.ReadClient.GetUser(context.Background(), &readPb.GetUserRequest{Email: req.GetEmail()})
	if err != nil {
		log.Printf("Error in controller.Login(), rpc call ReadClient.GetUser(): %v", err)
		return &pb.LoginResponse{}
	}
	if res.GetUser() == nil {
		log.Printf("Error in controller.Login(), Failed ReadClient.GetUser(): Can't find user with given email!")
		return &pb.LoginResponse{}
	}

	var response = &pb.LoginResponse{}

	if passwordIsValid([]byte(req.GetPassword()), []byte(res.GetHash())) {
		loggedInUser := res.GetUser()
		generatedJWT := generateJWT(loggedInUser.GetEmail())
		response = &pb.LoginResponse{
			User: &pb.SavedUser{
				Id:        loggedInUser.GetId(),
				Username:  loggedInUser.GetUsername(),
				Email:     loggedInUser.GetEmail(),
				FirstName: loggedInUser.GetFirstName(),
				LastName:  loggedInUser.GetLastName(),
				Avatar:    loggedInUser.GetAvatar(),
				Bio:       loggedInUser.GetBio(),
				Facebook:  loggedInUser.GetFacebook(),
				Instagram: loggedInUser.GetInstagram(),
				Twitter:   loggedInUser.GetTwitter(),
				CreatedAt: loggedInUser.GetCreatedAt(),
				UpdatedAt: loggedInUser.GetUpdatedAt(),
			},
			Token: generatedJWT,
		}
	}

	// ***************************************************************************************************************

	log.Printf("controller.Login(): [Response]: %+v\n", response)
	return response
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error in controller.hashAndSalt(): ", err)
	}
	return string(hash)
}

func passwordIsValid(pwd []byte, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, pwd)
	if err != nil {
		return false
	}
	return true
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
	signedString, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Println("Error in controller.generateJWT(): ", err)
		return ""
	}
	return signedString
}
