package main

import (
	"log"
	"time"

	"github.com/Spuxy/Goflexify/cache"
	"github.com/Spuxy/Goflexify/controller"
	"github.com/Spuxy/Goflexify/database"
	"github.com/Spuxy/Goflexify/router"
	"github.com/Spuxy/Goflexify/service"
	"github.com/Spuxy/Goflexify/utils/reader"
	"github.com/gofiber/fiber/v2"
	cacheFiber "github.com/gofiber/fiber/v2/middleware/cache"
)

func main() {
	cfg := reader.CreateReader("properties.ini")
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}
	redisConnection := cache.RedisHandler("localhost:6379", "", 0)
	redis := cache.NewRedisCache(redisConnection)
	userRepository := database.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	controller := controller.NewController(userService, redis)
	router2 := router.NewBackendRouter()
	router2.Factory(newFiber(), controller)
}

func newFiber() (app *fiber.App) {
	app = fiber.New()
	app.Use(cacheFiber.New(cacheFiber.Config{
		Expiration:   60 * time.Hour,
		CacheControl: true,
	}))
	return
}
