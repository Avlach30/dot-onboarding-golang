package dto

type OtpValidateRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required,e164"`
	Otp         string `json:"otp" validate:"required"`
}
