package controller

import (
	"github.com/NyanDesuyo/flashcard-service/config"
	"github.com/NyanDesuyo/flashcard-service/dto"
	"github.com/NyanDesuyo/flashcard-service/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
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

func Read(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["id"].(float64))

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("pageSize", "10"))

	if limit > 50 {
		limit = 50
	}

	offset := (page - 1) * limit

	var flashCards []models.FlashCard
	result := config.MainDB.Where("user_id = ?", userID).Limit(limit).Offset(offset).Find(&flashCards)
	if result.Error != nil {
		response := dto.GeneralResponseError{
			Message: "Cannot Get FlashCards",
			Error:   result.Error.Error(),
		}

		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	var totalData int64
	config.MainDB.Model(&models.FlashCard{}).Where("user_id = ?", userID).Count(&totalData)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success Get FlashCards",
		"data":    flashCards,
		"meta": fiber.Map{
			"total":     totalData,
			"page":      page,
			"limit":     limit,
			"last_page": (totalData + int64(limit) - 1) / int64(limit),
		},
	})
}
