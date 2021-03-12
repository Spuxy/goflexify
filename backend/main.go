package main

import (
	"log"

	"github.com/Spuxy/Goflexify/controller"
	"github.com/Spuxy/Goflexify/database"
	"github.com/Spuxy/Goflexify/router"
	"github.com/Spuxy/Goflexify/service"
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
	userService := service.NewUserService(userRepository)
	controller := controller.NewController(userService)
	app := fiber.New()
	router2 := router.NewBackendRouter()
	router2.Factory(app, controller)
}
