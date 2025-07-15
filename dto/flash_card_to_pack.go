package dto

type AddToPackRequest struct {
	PackID      uint `json:"pack_id"`
	FlashcardID uint `json:"flashcard_id"`
}

type AddToPackResponse struct {
	Message string `json:"message"`
}
