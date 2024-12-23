package dto

import "mime/multipart"

type UploadFilesRequest struct {
	Files []multipart.FileHeader `form:"files"` // Map of field names to file parts
}
