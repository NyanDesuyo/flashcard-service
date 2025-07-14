package dto

type GeneralResponseError struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}
