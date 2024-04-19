package dto

type AuthenticateInput struct {
	Key     string `json:"key" binding:"required"`
	PassKey string `json:"passKey" binding:"required"`
}
