package domain

import (
	"github.com/gin-gonic/gin"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/storage/dto"
)

type FileUsecase interface {
	UploadFiles(httpContext *gin.Context, files *dto.UploadFilesRequest) (map[string]string, error)
	Delete(httpContext *gin.Context, filePaths []string) error
	GetPresignUrlForUpload(httpContext *gin.Context, filePath string) (string, error)
	GetPresignUrlForDownload(httpContext *gin.Context, filePath string) (string, error)
	Get(httpContext *gin.Context, filePath string) (fileContents []byte, err error)
}
