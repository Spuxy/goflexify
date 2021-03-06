package router

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func Factory(r *fiber.App) {

	r.Get("test", func(c *fiber.Ctx) error {
		return c.SendString(os.Getenv("DB_DB"))
	})

	r.Listen(":5005")

}
