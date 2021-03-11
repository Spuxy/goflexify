package database

// bude potreba poresit !

import (
	"github.com/Spuxy/Goflexify/model"
)

type Repository interface {
	SelectUser(u *model.User) model.Users
	InsertUser(u *model.User) error
	UpdateUser(u *model.User) error
	DeleteUser(u *model.User) error
}
