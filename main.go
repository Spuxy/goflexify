package main

import (
	"fmt"
	"log"

	"github.com/Spuxy/Goflexify/controller"
	"github.com/Spuxy/Goflexify/database"
	"github.com/Spuxy/Goflexify/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	dbHandler := database.NewHandler(db)
	controller := controller.NewController(dbHandler)
	app := fiber.New()
	fmt.Printf("%T", app)
	router.Factory(app, controller)

}
