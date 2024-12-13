package dto

type AuthSignInResponse struct {
	Token string `json:"token" binding:"required"`
	Type  string `json:"type" binding:"required"`
}
