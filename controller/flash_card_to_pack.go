package controller

import (
	"github.com/NyanDesuyo/flashcard-service/config"
	"github.com/NyanDesuyo/flashcard-service/dto"
	"github.com/NyanDesuyo/flashcard-service/models"
	"github.com/gofiber/fiber/v2"
)

func AddFlashcardToPack(c *fiber.Ctx) error {
	req := new(dto.AddToPackRequest)
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

	var pack models.Pack
	if err := config.MainDB.First(&pack, req.PackID).Error; err != nil {
		response := dto.GeneralResponseError{
			Message: "Cannot find Pack",
			Error:   err.Error(),
		}

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	var flashcard models.FlashCard
	if err := config.MainDB.First(&flashcard, req.FlashcardID).Error; err != nil {
		response := dto.GeneralResponseError{
			Message: "Cannot find FlashCard",
			Error:   err.Error(),
		}

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if err := config.MainDB.Model(&pack).Association("FlashCards").Append(&flashcard); err != nil {
		response := dto.GeneralResponseError{
			Message: "Cannot find FlashCard",
			Error:   err.Error(),
		}

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := dto.AddToPackResponse{
		Message: "Success add flashcard to pack",
	}

	return c.Status(fiber.StatusCreated).JSON(response)

}
