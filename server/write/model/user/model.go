package user

import (
	"database/sql"
	_ "database/sql"
	"log"

	"github.com/fatih/structs"
	"github.com/nvhai245/cyberblog/server/write/connection"
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

// Insert func
func Insert(user *User) (success bool, userID int32) {
	rows, err := connection.DB.NamedQuery(`INSERT INTO users (username, email, hash, first_name, last_name, avatar, birthday, bio, facebook, instagram, twitter, created_at) VALUES (:Username, :Email, :Hash, :FirstName, :LastName, :Avatar, :Birthday, :Bio, :Facebook, :Instagram, :Twitter, :CreatedAt) RETURNING id`, structs.Map(user))
	if err != nil {
		log.Println("Error in userModel.Insert(): ", err)
		return false, 0
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&userID)
		if err != nil {
			log.Println("Error in userModel.Insert(): rows.Scan()", err)
			return false, 0
		}
	}
	return true, userID
}

// Update func
func Update(requestorEmail string, requestorIsAdmin bool, userToEdit *User) (success bool, user *User) {
	updatedUser := User{}
	var queryString string
	if requestorIsAdmin {
		queryString = `UPDATE users SET (username, first_name, last_name, avatar, birthday, bio, facebook, instagram, twitter, updated_at) = (:Username, :FirstName, :LastName, :Avatar, :Birthday, :Bio, :Facebook, :Instagram, :Twitter, :UpdatedAt) WHERE id = :ID RETURNING *`
	} else {
		queryString = `UPDATE users SET (username, first_name, last_name, avatar, birthday, bio, facebook, instagram, twitter, updated_at) = (:Username, :FirstName, :LastName, :Avatar, :Birthday, :Bio, :Facebook, :Instagram, :Twitter, :UpdatedAt) WHERE id = :ID AND email = :Email RETURNING *`
	}
	rows, err := connection.DB.NamedQuery(queryString, structs.Map(userToEdit))
	if err != nil {
		log.Println("Error in userModel.Update(): ", err)
		return false, nil
	}
	defer rows.Close()
	for rows.Next() {
		if rows.Err() == sql.ErrNoRows {
			log.Println("Error in userModel.Update(): no rows returned!", err)
			return false, nil
		}
		err = rows.StructScan(&updatedUser)
		if err != nil {
			log.Println("Error in userModel.Update(): rows.Scan()", err)
			return false, nil
		}
	}
	if structs.IsZero(updatedUser) {
		return false, nil
	}
	return true, &updatedUser
}
