package dto

type ProjectHistoryRes struct {
	HistoryType   string `json:"history_type"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	AttachmentUrl string `json:"attachment_url"`
	CreatedAt     string `json:"created_at"`
}
