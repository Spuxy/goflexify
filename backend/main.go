package main

import (
	"fmt"
	"log"

	"github.com/Spuxy/Goflexify/database"
	"github.com/Spuxy/Goflexify/model"
	"github.com/Spuxy/Goflexify/router"
	"github.com/gofiber/fiber/v2"
)

func main() {

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	dbHandler := database.NewHandler(db)
	_ = dbHandler.SelectUser(&model.User{})
	app := fiber.New()
	fmt.Printf("%T", app)
	router.Factory(app)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

}
