package model

type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password"  binding:"required"`
	RePassword string `json:"repassword"  binding:"required"`
}

type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
