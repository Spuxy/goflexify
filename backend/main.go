package main

import (
	"log"

	"github.com/Spuxy/Goflexify/controller"
	"github.com/Spuxy/Goflexify/database"
	"github.com/Spuxy/Goflexify/router"
	"github.com/Spuxy/Goflexify/utils/reader"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := reader.CreateReader("properties.ini")
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}
	userRepository := database.NewUserRepository(db)
	controller := controller.NewController(userRepository)
	app := fiber.New()
	router.Factory(app, controller)
}
