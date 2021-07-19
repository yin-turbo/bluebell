package model

type User struct {
	UserID   int64  `db : "user_Id"`
	Username string `db : "username"`
	Password string `db : "password"`
}
