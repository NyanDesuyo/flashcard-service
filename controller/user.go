package controller

import (
	"github.com/NyanDesuyo/flashcard-service/config"
	"github.com/NyanDesuyo/flashcard-service/dto"
	"github.com/NyanDesuyo/flashcard-service/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

func Register(c *fiber.Ctx) error {
	req := new(dto.CreateUserRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": "Cannot parse JSON",
				"error":   err.Error(),
			})
	}

	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": "Validation Error",
				"error":   err.Error(),
			})
	}

	hashedPassowrd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot hash password",
			"error":   err.Error(),
		})
	}

	newUser := models.User{
		UUID:     uuid.New(),
		Username: req.Username,
		Password: string(hashedPassowrd),
	}

	if result := config.MainDB.Create(&newUser); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot create user",
			"error":   result.Error.Error(),
		})
	}

	response := dto.CreateUserResponse{
		ID:       newUser.ID,
		UUID:     newUser.UUID.String(),
		Username: newUser.Username,
	}

	return c.Status(fiber.StatusOK).JSON(response)

}
