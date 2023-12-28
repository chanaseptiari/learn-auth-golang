package models

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LoginSearch struct {
	Username string
	Password string
}
type ResponseLogin struct {
	Token string
}
