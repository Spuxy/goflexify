package service

import (
	"errors"
	"regexp"

	"github.com/Spuxy/Goflexify/database"
	"github.com/Spuxy/Goflexify/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func NewUserService(repo database.UserRepositer) *UserService {
	return &UserService{repo}
}

type UserService struct {
	Repository database.UserRepositer
}

func (u *UserService) CheckAuthorization(cookie string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRETKEY), nil
	})

	return token, err
}

func (u *UserService) Create(user model.User) error {
	return nil
}

// return u.Repository.CreateUser(user)
func (u *UserService) GetUser(email, password string) (model.User, error) {
	if len(password) < 4 {
		return model.User{}, errors.New("your password is too weak")
	}
	if ok := emailRegex.MatchString(email); !ok {
		return model.User{}, errors.New("yout email has bad format")
	}
	return u.Repository.GetUserByEmailPassword(email, password)
}

func (u *UserService) Error(status int, msg string) error {
	return fiber.NewError(fiber.StatusBadRequest, msg)
}

func (u *UserService) GetKey() string {
	return SECRETKEY
}
