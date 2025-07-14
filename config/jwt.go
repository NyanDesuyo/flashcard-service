package config

import (
	"github.com/NyanDesuyo/flashcard-service/dto"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
)

func JWTMiddleware() fiber.Handler {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable not set")
	}

	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(jwtSecret)},
		ErrorHandler: jwtErrorHandler,
	})
}

func jwtErrorHandler(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		result := dto.GeneralResponseError{
			Message: "Missing or malformed JWT",
			Error:   err.Error(),
		}

		return c.Status(fiber.StatusBadRequest).JSON(result)
	}

	result := dto.GeneralResponseError{
		Message: "Invalid or expired JWT",
		Error:   err.Error(),
	}

	return c.Status(fiber.StatusUnauthorized).JSON(result)
}
