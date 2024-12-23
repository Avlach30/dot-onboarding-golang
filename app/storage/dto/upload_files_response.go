package dto

type UploadFilesResponse struct {
	FilePaths []UploadedFilePath `json:"file_paths"` // Map of field names to file parts
}

type UploadedFilePath struct {
	OriginalFile string `json:"original_file"`
	FilePath     string `json:"url"`
}
