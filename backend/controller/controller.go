package controller

import (
	"log"
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
	DB *database.DbHandler
}

var TokenName string = "AccessToken"

func NewController(db *database.DbHandler) *Controller {
	return &Controller{db}
}
func (c *Controller) List(http *fiber.Ctx) error {
	return http.SendString("HELLO WORLD")
}

func (c *Controller) Register(http *fiber.Ctx) error {
	time.Sleep(time.Second * 20)
	var user model.User
	log.Println(http.Body())
	http.BodyParser(&user)
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		fiber.NewError(fiber.StatusBadRequest, "Could not hash yout password :(")
	}

	user.Password = password
	rslt := c.DB.InsertUser(user)
	if rslt != nil {
		fiber.NewError(fiber.StatusBadRequest, "Something went wrong 😢")
	}

	http.SendStatus(fiber.StatusOK)
	return http.JSON(fiber.Map{"message": "Account has been created 😂"})
}

func (c *Controller) Login(http *fiber.Ctx) (err error) {
	var user model.User

	err = http.BodyParser(&user)
	if err != nil {
		fiber.NewError(fiber.StatusBadRequest, "Something went wrong 😢")
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := claims.SignedString([]byte(SECRETKEY))
	if err != nil {
		fiber.NewError(fiber.StatusBadRequest, "Something went wrong 😢")
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

func (c *Controller) CreateRoom(http *fiber.Ctx) error {
	tkn := http.Cookies("AccessToken")
	log.Println(tkn)
	return nil
}

// LIP UDELAT VYHAZOVANI CHYB !
