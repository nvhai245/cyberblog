package model

import (
	_ "database/sql"
	"log"

	"github.com/fatih/structs"
	"github.com/nvhai245/cyberblog/server/write/connection"
)

type NewUser struct {
	ID        string `db:"id"`
	Username  string `db:"username"`
	Email     string `db:"email"`
	Hash      string `db:"hash"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Avatar    string `db:"avatar"`
	Birthday  int32  `db:"birthday"`
	Bio       string `db:"bio"`
	Facebook  string `db:"facebook"`
	Instagram string `db:"instagram"`
	Twitter   string `db:"twitter"`
}

// Insert func
func Insert(user *NewUser) (success bool) {
	_, err := connection.DB.NamedExec(`INSERT INTO users (username, email, hash, first_name, last_name, avatar, birthday, bio, facebook, instagram, twitter) VALUES (:Username, :Email, :Hash, :FirstName, :LastName, :Avatar, :Birthday, :Bio, :Facebook, :Instagram, :Twitter)`, structs.Map(user))
	if err != nil {
		log.Println("Error in model.Insert(): ", err)
		return false
	}
	return true
}
