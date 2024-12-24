package handler

import (
	"bytes"
	"io"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/storage/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/storage/dto"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/middleware"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/singleton"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"

	"github.com/gin-gonic/gin"
)

type CommonHandler struct {
	fileUsecase domain.FileUsecase
}

func NewCommonHandler(router *gin.Engine, fileUsecase domain.FileUsecase) {
	commonRoute := router.Group("/api/v1/commons")

	common := &CommonHandler{
		fileUsecase: fileUsecase,
	}

	commonRoute.POST("/upload-files", middleware.ValidateRequestFormData(&dto.UploadFilesRequest{}), common.CreateFile())
	commonRoute.GET("/download-file", common.DownloadFile())
	commonRoute.POST("/generate-presign-url", middleware.ValidateRequestJSON(&dto.CreatePresignURLUploadRequest{}), common.GeneratePresignURLUpload())
	commonRoute.DELETE("/delete-files", middleware.ValidateRequestJSON(&dto.CreatePresignURLUploadRequest{}), common.GeneratePresignURLUpload())
}

func (common *CommonHandler) CreateFile() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		createFileRequest := singleton.GetHTTPRequest[dto.UploadFilesRequest](httpContext)
		filePaths, _ := common.fileUsecase.UploadFiles(httpContext, createFileRequest)

		response := &dto.UploadFilesResponse{
			FilePaths: make([]dto.UploadedFilePath, len(filePaths)),
		}

		i := 0
		for originalFile, filePath := range filePaths {
			response.FilePaths[i] = dto.UploadedFilePath{
				OriginalFile: originalFile,
				FilePath:     filePath,
			}
			i++
		}

		httpContext.JSON(http.StatusOK, utils.SucessResponse(
			&response,
		))
	}
}

func (common *CommonHandler) DeleteFile() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		filePaths := strings.Split(httpContext.Query("file_paths"), ",")

		common.fileUsecase.Delete(httpContext, filePaths)
		httpContext.JSON(http.StatusOK, utils.SucessResponse(
			nil,
		))
	}
}

func (common *CommonHandler) GeneratePresignURLUpload() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		GeneratePresignURL := singleton.GetHTTPRequest[dto.CreatePresignURLUploadRequest](httpContext)
		var url string
		switch GeneratePresignURL.ActionType {
		case "download":
			url, _ = common.fileUsecase.GetPresignUrlForDownload(httpContext, GeneratePresignURL.FileName)
		case "upload":
			url, _ = common.fileUsecase.GetPresignUrlForUpload(httpContext, GeneratePresignURL.FileName)
		default:
			panic(*exception.BussinessException("action_type did not recognize"))
		}

		httpContext.JSON(http.StatusOK, utils.SucessResponse(
			&dto.CreatePresignURLUploadResponse{
				PresignURL: url,
			},
		))
	}
}

func (common *CommonHandler) DownloadFile() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		filePath := httpContext.Query("file_path")
		if filePath == "" {
			panic(*exception.BussinessException("file_path is required"))
		}

		// Get the file reader from the usecase
		fileContents, err := common.fileUsecase.Get(httpContext, filePath)
		if err != nil {
			panic(*exception.BussinessException("Failed to get file contents"))
		}

		fileReader := bytes.NewReader(fileContents)
		fileSize := int64(len(fileContents))

		// Set headers for file download
		httpContext.Header("Content-Description", "File Transfer")
		httpContext.Header("Content-Transfer-Encoding", "binary")
		httpContext.Header("Content-Disposition", "attachment; filename="+filepath.Base(filePath))
		httpContext.Header("Content-Type", "application/octet-stream")
		httpContext.Header("Content-Length", strconv.FormatInt(fileSize, 10))

		// Stream the file content
		httpContext.Stream(func(w io.Writer) bool {
			_, err := io.Copy(w, fileReader)
			return err == nil
		})
	}
}
