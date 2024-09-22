package dto

type GetProfileResponse struct {
	Fullname    string `json:"fullname" `
	ImageURL    string `json:"image_url" `
	Email       string `json:"email" `
	PhoneNumber string `json:"phone_number" `
	Role        string `json:"role,omitempty" `
}

type GetProfileTalentResponse struct {
	Fullname string `json:"fullname" `
	ImageURL string `json:"image_url" `
	Role     string `json:"role,omitempty" `
}
