package router

import (
	"log"

	"github.com/Spuxy/Goflexify/controller"
	"github.com/gofiber/fiber/v2"
)

type Routerable interface {
	Factory(r *fiber.App, c *controller.Controller)
}

type BackendRouter struct{}

func NewBackendRouter() Routerable {
	return &BackendRouter{}
}

type FrontendRouter struct{}

func NewFrontendRouter() Routerable {
	return &FrontendRouter{}
}

func (b *FrontendRouter) Factory(r *fiber.App, c *controller.Controller) {
	// [#] GROUPS [#]
	api := r.Group("/api")

	// [#] VERSIONS [#]
	v1 := api.Group("/v1")

	// [#] USERS [#]
	v1.Get("/user/:id", c.Info)
	v1.Post("/login", c.Login)
	v1.Post("/logout", c.Logout)
	v1.Post("/user", c.Register)
	err := r.Listen(":5007")
	if err != nil {
		log.Fatal("Something went wront with httpServer", err)
	}

}

func (b *BackendRouter) Factory(r *fiber.App, c *controller.Controller) {
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
