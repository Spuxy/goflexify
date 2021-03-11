package controller

import (
	"strconv"
	"time"

	"github.com/Spuxy/Goflexify/database"
	"github.com/Spuxy/Goflexify/model"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const SECRETKEY string = "siliconvalley"

type Controller struct {
	UserRepository database.IUserRepository
}

var TokenName string = "AccessToken"

func NewController(db database.IUserRepository) *Controller {
	return &Controller{db}
}

func (c *Controller) Info(http *fiber.Ctx) error {
	cookie := http.Cookies(TokenName)
	token, err := checkAuthorization(cookie)
	if err != nil {
		http.Status(fiber.StatusUnauthorized)
		return http.JSON(fiber.Map{
			"message": "unauthorized",
		})
	}
	claims := token.Claims
	_ = http.Params("id")
	return http.JSON(claims)
}

func (c *Controller) Register(http *fiber.Ctx) error {
	var user model.User
	http.BodyParser(&user)
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		fiber.NewError(fiber.StatusBadRequest, "Could not hash yout password :(")
	}

	user.Password = string(password)
	rslt := c.UserRepository.CreateUser(user)
	if rslt != nil {
		fiber.NewError(fiber.StatusForbidden, "could not create the account 😢")
	}

	http.SendStatus(fiber.StatusOK)
	return http.JSON(fiber.Map{"message": "Account has been created 😂"})
}

func (c *Controller) Login(http *fiber.Ctx) (err error) {
	var user model.User
	err = http.BodyParser(&user)
	if err != nil {
		fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	foundUser, err := c.UserRepository.GetUserByEmailPassword(user.Email, user.Password)
	if err != nil {
		fiber.NewError(fiber.StatusBadRequest, "Something went wrong 😢")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password)); err != nil {
		fiber.NewError(fiber.StatusForbidden, "could not hashed your password 😢")
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := claims.SignedString([]byte(SECRETKEY))
	if err != nil {
		fiber.NewError(fiber.StatusBadRequest, "could not encrypt the key of jwt 😢")
		return
	}
	http.Cookie(&fiber.Cookie{
		Name:     TokenName,
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	})
	return http.JSON(fiber.Map{
		"message": "success",
	})
}

func (c *Controller) Logout(r *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     TokenName,
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	r.Cookie(&cookie)
	return r.JSON(fiber.Map{
		"message": "success",
	})
}

func checkAuthorization(cookie string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRETKEY), nil
	})
	return token, err
}

// LIP UDELAT VYHAZOVANI CHYB !
