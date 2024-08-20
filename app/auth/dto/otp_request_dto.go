package dto

type OtpRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required,e164"`
}
