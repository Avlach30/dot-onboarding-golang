package dto

type ExchangeRequest struct {
	PhoneNumber     string `json:"phone_number" validate:"required,e164"`
	FirebaseIdToken string `json:"firebase_id_token" validate:"required"`
}

type ExchangeResponse struct {
	Token string `json:"token" `
}
