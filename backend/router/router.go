package router

import (
	"log"

	"github.com/Spuxy/Goflexify/controller"
	"github.com/gofiber/fiber/v2"
)

func Factory(r *fiber.App, c *controller.Controller) {
	// [#] GROUPS [#]
	api := r.Group("/api")

	// [#] VERSIONS [#]
	v1 := api.Group("/v1")

	// [#] USERS [#]
	v1.Get("/user/:id", c.Info)
	v1.Post("/login", c.Login)
	v1.Post("/logout", c.Logout)
	v1.Post("/user", c.Register)
	err := r.Listen(":5005")
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
