package controller

import (
	"github.com/NyanDesuyo/flashcard-service/config"
	"github.com/NyanDesuyo/flashcard-service/dto"
	"github.com/NyanDesuyo/flashcard-service/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

func Register(c *fiber.Ctx) error {
	req := new(dto.UserRequest)
	if err := c.BodyParser(req); err != nil {
		response := dto.GeneralResponseError{
			Message: "Cannot Parse Body",
			Error:   err.Error(),
		}

		return c.Status(fiber.StatusBadRequest).JSON(
			response)
	}

	if err := validate.Struct(req); err != nil {
		response := dto.GeneralResponseError{
			Message: "Validation Error",
			Error:   err.Error(),
		}

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		response := dto.GeneralResponseError{
			Message: "Cannot has password",
			Error:   err.Error(),
		}

		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	newUser := models.User{
		UUID:     uuid.New(),
		Username: req.Username,
		Password: string(hashedPassword),
	}

	if result := config.MainDB.Create(&newUser); result.Error != nil {
		response := dto.GeneralResponseError{
			Message: "Cannot create user",
			Error:   result.Error.Error(),
		}

		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := dto.CreateUserResponse{
		Message:  "User Created",
		ID:       newUser.ID,
		UUID:     newUser.UUID.String(),
		Username: newUser.Username,
	}

	return c.Status(fiber.StatusOK).JSON(response)

}

func Login(c *fiber.Ctx) error {
	req := new(dto.UserRequest)
	if err := c.BodyParser(req); err != nil {
		response := dto.GeneralResponseError{
			Message: "Cannot parse Body",
			Error:   err.Error(),
		}

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if err := validate.Struct(req); err != nil {
		response := dto.GeneralResponseError{
			Message: "Validation Error",
			Error:   err.Error(),
		}

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	var user models.User
	if result := config.MainDB.Where("username = ?", req.Username).First(&user); result.Error != nil {
		response := dto.GeneralResponseError{
			Message: "Cannot find user",
			Error:   result.Error.Error(),
		}

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		response := dto.GeneralResponseError{
			Message: "Wrong password",
			Error:   err.Error(),
		}

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	claims := jwt.MapClaims{
		"username": user.Username,
		"uuid":     user.UUID.String(),
		"id":       user.ID,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := os.Getenv("JWT_SECRET")

	t, err := token.SignedString([]byte(secret))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot sign JWT",
			"error":   err.Error(),
		})
	}

	response := dto.ReadUserTokenResponse{
		Message:  "User login success",
		Username: user.Username,
		Token:    t,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
