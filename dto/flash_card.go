package dto

type CreateFlashCardRequest struct {
	Front string `json:"front"`
	Back  string `json:"back"`
}

type ReadFlashCardResponse struct {
	Message string `json:"message"`
	Front   string `json:"front"`
	Back    string `json:"back"`
}
