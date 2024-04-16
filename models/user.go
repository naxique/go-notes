package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUser(username, password string) *User {
	return &User{
		Username: username,
		Password: password,
	}
}
