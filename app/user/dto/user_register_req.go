package dto

type RegisterRequest struct {
	Fullname    string `json:"fullname" validate:"required"`
	Email       string `json:"email" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required,e164"`
	Type        string `json:"type"`
}
