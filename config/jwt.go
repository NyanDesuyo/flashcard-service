package config

import (
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
		SigningKey: jwtware.SigningKey{Key: []byte(jwtSecret)},
	})
}
