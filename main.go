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

	user := app.Group("/user")
	user.Post("/register", controller.Register)
	user.Post("/login", controller.Login)

	app.Use(config.JWTMiddleware())

	app.Get("/protected", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "this is protected endpoint",
		})
	})

	flashcard := app.Group("/flashcard")
	flashcard.Post("/create", controller.Create)
	flashcard.Get("/read", controller.Read)

	app.Listen(":8080")
}
