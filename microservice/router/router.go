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
	r.Get("/login", c.Login)
	r.Post("/user", c.Register)
	r.Get("/micro", func(c *fiber.Ctx) error {
		log.Println("h")
		//Create a folder/directory at a full qualified path
		err := os.Mkdir("/home/boyekofif/programming/go/testujuto", 0755)
		if err != nil {
			log.Fatal(err)
		}
		return c.SendString("hgg")
	})
	err := r.Listen(":5006")
	if err != nil {
		log.Fatal("Something went wront with httpServer", err)
	}

}

func test(t int64) func(c *fiber.Ctx) error {
	var A []int64
	A = append(A, t, 323)
	return func(c *fiber.Ctx) error {
		return c.JSON(A)
	}
}
