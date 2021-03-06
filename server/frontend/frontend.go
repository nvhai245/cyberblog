package main

import (
	"flag"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/nvhai245/cyberblog/server/frontend/middleware"
	readPb "github.com/nvhai245/cyberblog/server/read/proto"
	writePb "github.com/nvhai245/cyberblog/server/write/proto"
	"log"
	"net/http"
	"os"

	authPb "github.com/nvhai245/cyberblog/server/auth/proto"
	"github.com/nvhai245/cyberblog/server/frontend/connection"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/joho/godotenv/autoload"
	"github.com/nvhai245/cyberblog/server/frontend/graph"
	"github.com/nvhai245/cyberblog/server/frontend/graph/generated"
	"google.golang.org/grpc"
)

const defaultPort = "8080"

func main() {
	flag.Parse()
	_ = godotenv.Load("cyberblog.env")
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

	// This middleware injects the ResponseWriter and Reader structs
	// into the context of each resolver so they have access http headers and cookies. The
	// session is also passed in so that resolvers will ultimately have access to the store.
	handlerWithMiddleware := middleware.InjectHTTPMiddleware(store)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", handlerWithMiddleware(srv))

	// ---gRPC Dials---
	// Auth
	authConn, err := grpc.Dial(os.Getenv("AUTH_SERVICE_HOST"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Connected to Auth service")
	}
	defer authConn.Close()
	connection.AuthClient = authPb.NewAuthClient(authConn)

	// Read
	readConn, err := grpc.Dial(os.Getenv("READ_SERVICE_HOST"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Connected to Read service")
	}
	defer readConn.Close()
	connection.ReadClient = readPb.NewReadClient(readConn)

	// Write
	writeConn, err := grpc.Dial(os.Getenv("WRITE_SERVICE_HOST"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Connected to Write service")
	}
	defer readConn.Close()
	connection.WriteClient = writePb.NewWriteClient(writeConn)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
