package dto

type ProjectHistoryRes struct {
	HistoryType     string `json:"history_type"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	AttachmentUrl   string `json:"attachment_url"`
	AttachmentTitle string `json:"attachment_title"`
	CreatedAt       string `json:"created_at"`
}
