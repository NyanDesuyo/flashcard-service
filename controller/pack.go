package controller

import (
	"github.com/NyanDesuyo/flashcard-service/config"
	"github.com/NyanDesuyo/flashcard-service/dto"
	"github.com/NyanDesuyo/flashcard-service/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func CreatePack(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["id"].(float64))

	req := new(dto.CreatePackRequest)
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

	newPack := models.Pack{
		Name:   req.Name,
		UserID: userID,
	}

	if err := config.MainDB.Create(&newPack).Error; err != nil {
		response := dto.GeneralResponseError{
			Message: "Cannot Create Pack",
			Error:   err.Error(),
		}

		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := dto.CreatePackResponse{
		Message: "Success create pack",
		ID:      newPack.ID,
		Name:    newPack.Name,
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
