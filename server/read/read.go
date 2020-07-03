package main

import (
	_ "database/sql"
	"flag"
	"log"
	"net"
	"os"

	"github.com/nvhai245/cyberblog/server/read/connection"
	pb "github.com/nvhai245/cyberblog/server/read/proto"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	_ = godotenv.Load("cyberblog.env")
	var err error

	connection.DB, err = sqlx.Connect("postgres", os.Getenv("POSTGRES_STRING"))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Successfully connected to postgres DB!")

	lis, err := net.Listen("tcp", os.Getenv("READ_SERVICE_HOST"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Println("Read service listening on port 8082...")
	}
	grpcServer := grpc.NewServer()
	pb.RegisterReadServer(grpcServer, &pb.UnimplementedReadServer{})
	// determine whether to use TLS
	grpcServer.Serve(lis)
}
