package model

import (
	"database/sql"
	"github.com/nvhai245/cyberblog/server/read/connection"
	"log"
)

type User struct {
	ID        int32  `db:"id"`
	Username  string `db:"username"`
	Email     string `db:"email"`
	Hash      string `db:"hash"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Avatar    string `db:"avatar"`
	Birthday  int64  `db:"birthday"`
	Bio       string `db:"bio"`
	Facebook  string `db:"facebook"`
	Instagram string `db:"instagram"`
	Twitter   string `db:"twitter"`
	IsAdmin   bool   `db:"is_admin"`
	CreatedAt int64  `db:"created_at"`
	UpdatedAt int64  `db:"updated_at"`
}

// GetUserByEmail func
func GetUserByEmail(email string) *User {
	foundUser := User{}
	rows, err := connection.DB.Queryx(`SELECT * FROM users WHERE email = $1`, email)
	if err != nil {
		log.Println("Error in model.GetUserByEmail(): ", err)
		return nil
	}
	if rows.Err() == sql.ErrNoRows {
		log.Println("Error in model.GetUserByEmail(): ", err)
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.StructScan(&foundUser)
		if err != nil {
			log.Println("Error in model.GetUserByEmail(): rows.StructScan()", err)
			return nil
		}
	}
	log.Printf("%+v", foundUser)
	return &foundUser
}
