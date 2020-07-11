package connection

import (
	_ "database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var (
	// DB to use globally
	DB *sqlx.DB
)

func init() {
	_ = godotenv.Load("cyberblog.env")
	var err error
	DB, err = sqlx.Connect("postgres", os.Getenv("POSTGRES_STRING"))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Successfully connected to postgres DB!")
}
