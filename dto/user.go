package dto

type UserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type CreateUserResponse struct {
	Message  string `json:"message"`
	ID       uint   `json:"id"`
	UUID     string `json:"uuid"`
	Username string `json:"username"`
}

type ReadUserTokenResponse struct {
	Message  string `json:"message"`
	Token    string `json:"token"`
	Username string `json:"username"`
}
