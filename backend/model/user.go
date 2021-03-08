package model

type Users []User

type User struct {
	ID       uint
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      uint   `json:"age"`
	Password string `json:"password" form:"password"`
}
