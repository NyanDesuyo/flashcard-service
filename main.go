package main

import (
	"github.com/NyanDesuyo/flashcard-service/config"
	"github.com/NyanDesuyo/flashcard-service/controller"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	config.ConnectMainPostgres()
	config.LoadEnv()

	app := fiber.New(fiber.Config{
		AppName: "FlashCard Services 1.0.0",
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello World!",
		})
	})

	app.Post(
		"/user/register",
		controller.Register,
	)

	app.Use(config.JWTMiddleware())

	app.Listen(":8080")
}
