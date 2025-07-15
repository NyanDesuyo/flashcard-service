package dto

type CreatePackRequest struct {
	Name string `json:"name" binding:"required"`
}

type CreatePackResponse struct {
	Message string `json:"message"`
	ID      uint   `json:"id"`
	Name    string `json:"name"`
}
