package main

import (
	"log"
	"net/http"
	"os"
	"flag"

	"github.com/nvhai245/cyberblog/server/frontend/connection"
	authPb "github.com/nvhai245/cyberblog/server/auth/proto"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/nvhai245/cyberblog/server/frontend/graph"
	"github.com/nvhai245/cyberblog/server/frontend/graph/generated"
	"google.golang.org/grpc"
)

const defaultPort = "8080"

func main() {
	flag.Parse()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	// ---gRPC Dials---
	// Auth
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Connected to Auth service")
	}
	defer conn.Close()
	connection.AuthClient = authPb.NewAuthClient(conn)



	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
