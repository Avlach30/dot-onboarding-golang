package dto

type CreatePresignURLUploadResponse struct {
	PresignURL string `json:"presign_url"` // Map of field names to file parts
}
