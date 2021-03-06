package router

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func Factory(r *fiber.App) {

	r.Get("test", func(c *fiber.Ctx) error {
		return c.SendString(os.Getenv("DB_DB"))
	})

	err := r.Listen(":5005")

	if err != nil {
		log.Fatal("Something went wront with httpServer")
	}

}
