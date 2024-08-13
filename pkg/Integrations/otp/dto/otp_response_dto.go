package dto

type OtpResponse struct {
	MessageID string `json:"messageId"`
	To        string `json:"to"`
	Status    string `json:"status"`
	Text      string `json:"text"`
	Cost      string `json:"cost"`
}
