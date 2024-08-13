package dto

type OtpRequest struct {
	Userkey string `json:"userkey"`
	Passkey string `json:"passkey"`
	To      string `json:"to"`
	Brand   string `json:"brand"`
	Otp     string `json:"otp"`
}
