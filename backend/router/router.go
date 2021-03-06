package router

import (
	"log"
	"os"

	"github.com/Spuxy/Goflexify/controller"
	"github.com/gofiber/fiber/v2"
)

func Factory(r *fiber.App, c *controller.Controller) {
	r.Get("/user/:id", func(c *fiber.Ctx) error {
		return c.SendString(os.Getenv("DB_DB"))
	})
	r.Get("/users", c.List)
	r.Post("/user", c.Insert)
	err := r.Listen(":5005")

	if err != nil {
		log.Fatal("Something went wront with httpServer")
	}

}
