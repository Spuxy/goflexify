package controller

import (
	"github.com/Spuxy/Goflexify/database"
	"github.com/Spuxy/Goflexify/model"
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   uint8  `json:"age"`
}
type Controller struct {
	DB *database.DbHandler
}

func NewController(db *database.DbHandler) *Controller {
	return &Controller{db}
}
func (c *Controller) List(http *fiber.Ctx) error {
	return http.SendString("HELLO WORLD")
}
func (c *Controller) Insert(http *fiber.Ctx) error {
	var body map[string]string
	http.BodyParser(&body)
	rslt := c.DB.InsertUser(&model.User{
		Name:  body["name"],
		Email: body["email"],
		Age:   body["age"],
	})
	if rslt != nil {
		fiber.NewError(fiber.StatusBadRequest, "Something went wrong 😢")
	}
	http.SendStatus(fiber.StatusOK)
	return http.JSON(fiber.Map{"message": "Account has been created 😂"})
}
