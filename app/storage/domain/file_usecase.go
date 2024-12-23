package domain

import (
	"context"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/storage/dto"
)

type FileUsecase interface {
	UploadFiles(context *context.Context, files *dto.UploadFilesRequest) (map[string]string, error)
	Delete(context *context.Context, filePaths []string) error
	GetPresignUrlForUpload(context *context.Context, filePath string) (string, error)
	GetPresignUrlForDownload(context *context.Context, filePath string) (string, error)
	Get(context *context.Context, filePath string) (fileContents []byte, err error)
}
