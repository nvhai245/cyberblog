package main

import (
	"flag"
	"log"
	"net"
	"os"

	pb "github.com/nvhai245/cyberblog/server/write/proto"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", os.Getenv("WRITE_SERVICE_HOST"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Println("Write service listening on port 8083...")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterWriteServer(grpcServer, &WriteServer{})
	// determine whether to use TLS
	grpcServer.Serve(lis)
}
