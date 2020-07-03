package main

import (
	"flag"
	"log"
	"net"
	"os"

	pb "github.com/nvhai245/cyberblog/server/auth/proto"
	"github.com/nvhai245/cyberblog/server/auth/connection"
	writePb "github.com/nvhai245/cyberblog/server/write/proto"
	readPb "github.com/nvhai245/cyberblog/server/read/proto"

	"google.golang.org/grpc"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	flag.Parse()
	_ = godotenv.Load("cyberblog.env")

	// ---gRPC Dials---
	// Write
	writeConn, err := grpc.Dial(os.Getenv("WRITE_SERVICE_HOST"), grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Connected to WRITE service")
	}
	defer writeConn.Close()
	connection.WriteClient = writePb.NewWriteClient(writeConn)

	// Read
	readConn, err := grpc.Dial(os.Getenv("READ_SERVICE_HOST"), grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Connected to READ service")
	}
	defer readConn.Close()
	connection.ReadClient = readPb.NewReadClient(readConn)

	// gRPC Auth server
	lis, err := net.Listen("tcp", os.Getenv("AUTH_SERVICE_HOST"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Println("Auth service listening on port 9090...")
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAuthServer(grpcServer, &AuthServer{})
	// determine whether to use TLS
	grpcServer.Serve(lis)
}