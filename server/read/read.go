package main

import (
	_ "database/sql"
	"flag"
	"log"
	"net"
	"os"

	pb "github.com/nvhai245/cyberblog/server/read/proto"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", os.Getenv("READ_SERVICE_HOST"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Println("Read service listening on port 8082...")
	}
	grpcServer := grpc.NewServer()
	pb.RegisterReadServer(grpcServer, &ReadServer{})
	// determine whether to use TLS
	grpcServer.Serve(lis)
}
