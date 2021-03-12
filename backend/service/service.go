package service

import (
	"github.com/Spuxy/Goflexify/model"
	"github.com/dgrijalva/jwt-go"
)

const SECRETKEY string = "siliconvalley"

type Servicer interface {
	Create(user model.User) error
	CheckAuthorization(cookie string) (*jwt.Token, error)
	GetUser(email, password string) (model.User, error)
	Error(status int, msg string) error
	GetKey() string
}
