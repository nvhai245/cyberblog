package model

import (
	_ "database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/nvhai245/cyberblog/server/read/connection"
)

type NewUser struct {
	Username  string `db:"username"`
	Email     string `db:"email"`
	Hash      string `db:"hash"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Avatar    string `db:"avatar"`
	Birthday  string `db:"birthday"`
	Bio       string `db:"bio"`
	Facebook  string `db:"facebook"`
	Instagram string `db:"instagram"`
	Twitter   string `db:"twitter"`
}

// Insert func
func Insert(db *sqlx.DB, user NewUser) (success bool) {
	_, err := db.NamedExec("INSERT INTO users (username, email, hash, first_name, last_name, avatar, birthday, bio, facebook, instagram, twitter) VALUES (:username, :email, :hash, :first_name, :last_name, :avatar, :birthday, :bio, :facebook, :instagram, :twitter)", user)
	if err != nil {
		log.Println("Error in model.Insert(): ", err)
	}
}