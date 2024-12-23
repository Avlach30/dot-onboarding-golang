package dto

type CreatePresignURLUploadRequest struct {
	FileName   string `json:"file_name"`   // Map of field names to file parts
	ActionType string `json:"action_type"` // Map of field names to file parts
}
