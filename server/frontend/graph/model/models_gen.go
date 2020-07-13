// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AuthResponse struct {
	Message string `json:"message"`
	User    *User  `json:"user"`
}

type NewUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Token    string `json:"token"`
	ExpireAt int    `json:"expire_at"`
}

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
	Birthday  int    `json:"birthday"`
	Bio       string `json:"bio"`
	Facebook  string `json:"facebook"`
	Instagram string `json:"instagram"`
	Twitter   string `json:"twitter"`
	IsAdmin   bool   `json:"is_admin"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
}
