package dto

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=8,max=31"`
}

type CreateUserResponse struct {
	ID       uint   `json:"id"`
	UUID     string `json:"uuid"`
	Username string `json:"username"`
}
