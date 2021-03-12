package database

import (
	"github.com/Spuxy/Goflexify/model"
)

type UserRepositer interface {
	GetUserById(id int) (model.User, error)
	GetUserByEmailPassword(email, password string) (model.User, error)
	CreateUser(u *model.User) error
	UpdateUser(u *model.User) error
	DeleteUserById(id int) error
}
