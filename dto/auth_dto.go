package dto

type ExchangeRequest struct {
	FirebaseIdToken string `json:"firebase_id_token" validate:"required"`
}

type ExchangeResponse struct {
	Token string `json:"token" `
}

type RegisterRequest struct {
	Fullname string `json:"fullname" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Type     string `json:"type"`
}
