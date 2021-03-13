package controller

import (
	"strconv"
	"time"

	"github.com/Spuxy/Goflexify/cache"
	"github.com/Spuxy/Goflexify/model"
	"github.com/Spuxy/Goflexify/service"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type Controller struct {
	UserService service.Servicer
	Cache       cache.Cacher
}

var TokenName string = "AccessToken"

func NewController(service service.Servicer, cache cache.Cacher) *Controller {
	return &Controller{
		UserService: service,
		Cache:       cache,
	}
}

func (c *Controller) Info(http *fiber.Ctx) error {
	val, err := c.Cache.Get("hello")
	if err != nil {
		http.Status(fiber.StatusBadRequest)
		return http.JSON(fiber.Map{
			"message": "zkouska",
		})
	}
	return http.SendString(string(val))
	cookie := http.Cookies(TokenName)
	token, err := c.UserService.CheckAuthorization(cookie)
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

	err := http.BodyParser(&user)
	if err != nil {
		return c.UserService.Error(fiber.StatusBadRequest, "we could not process your request")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		c.UserService.Error(fiber.StatusForbidden, "Could not hash yout password :(")
	}

	user.Password = string(password)
	rslt := c.UserService.Create(user)
	if rslt != nil {
		c.UserService.Error(fiber.StatusForbidden, "could not create the account 😢")
	}

	http.SendStatus(fiber.StatusOK)
	return http.JSON(fiber.Map{"message": "Account has been created 😂"})
}

func (c *Controller) Login(http *fiber.Ctx) (err error) {
	var user model.User
	err = http.BodyParser(&user)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	foundUser, err := c.UserService.GetUser(user.Email, user.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password)); err != nil {
		return fiber.NewError(fiber.StatusForbidden, "could not hashed your password 😢")
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := claims.SignedString([]byte(c.UserService.GetKey()))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "could not encrypt the key of jwt 😢")
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

// LIP UDELAT VYHAZOVANI CHYB !
