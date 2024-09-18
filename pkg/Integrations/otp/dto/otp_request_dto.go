package dto

type OtpRequestWaOfficial struct {
	Userkey string `json:"userkey"`
	Passkey string `json:"passkey"`
	To      string `json:"to"`
	Brand   string `json:"brand"`
	Otp     string `json:"otp"`
}

type OtpRequestSmsMasking struct {
	Userkey string `json:"userkey"`
	Passkey string `json:"passkey"`
	To      string `json:"to"`
	Message string `json:"message"`
}
