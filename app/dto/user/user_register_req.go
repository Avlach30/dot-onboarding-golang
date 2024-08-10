package userdto

type RegisterRequest struct {
	Fullname    string `json:"fullname" validate:"required"`
	Email       string `json:"email" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Type        string `json:"type"`
}
