package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New(fiber.Config{
		AppName: "FlashCard Services 1.0.0",
	})

	app.Get("/hello/:name?", func(c *fiber.Ctx) error {
		if c.Params("name") != "" {
			return c.JSON(fiber.Map{
				"message": "Hello " + c.Params("name") + "!",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Hello, who are you?",
		})
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello World!",
		})
	})

	app.Listen(":8080")
}
