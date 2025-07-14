package controller

import (
	"github.com/NyanDesuyo/flashcard-service/config"
	"github.com/NyanDesuyo/flashcard-service/dto"
	"github.com/NyanDesuyo/flashcard-service/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Create(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["id"].(float64))

	req := new(dto.CreateFlashCardRequest)
	if err := c.BodyParser(req); err != nil {
		response := dto.GeneralResponseError{
			Message: "Cannot Parse Body",
			Error:   err.Error(),
		}

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if err := validate.Struct(req); err != nil {
		response := dto.GeneralResponseError{
			Message: "Invalid Request",
			Error:   err.Error(),
		}

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	newFlashCard := models.FlashCard{
		Front:  req.Front,
		Back:   req.Back,
		UserId: userID,
	}

	if err := config.MainDB.Create(&newFlashCard).Error; err != nil {
		response := dto.GeneralResponseError{
			Message: "Cannot Create FlashCard",
			Error:   err.Error(),
		}

		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := dto.ReadFlashCardResponse{
		Message: "Success create flashcard",
		Front:   newFlashCard.Front,
		Back:    newFlashCard.Back,
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
