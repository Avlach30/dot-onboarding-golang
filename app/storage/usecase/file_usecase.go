package usecase

import (
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/storage/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/storage/dto"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/config"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/singleton"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"
)

const FileUploads = "uploads"

// NewFileUsecase create new instance of FileUsecase
func NewFileUsecase() domain.FileUsecase {
	return &FileUsecase{}
}

// FileUsecase represent the file's usecase
type FileUsecase struct {
}

// Create implements domain.FileUsecase.
func (f *FileUsecase) UploadFiles(httpContext *gin.Context, files *dto.UploadFilesRequest) (map[string]string, error) {
	allowUploadSizeGeneral, err := strconv.ParseInt(config.MaxSizeUploadFile, 10, 64)
	if err != nil {
		panic(*exception.BussinessException("Wrong Max Upload Size"))
	}

	fileContents := make(map[string][]byte, len(files.Files))

	for _, file := range files.Files {
		fileContent, err := file.Open()
		if err != nil {
			panic(*exception.BussinessException("Cannot read the file"))
		}
		defer fileContent.Close()

		fileBytes, err := io.ReadAll(fileContent)
		if err != nil {
			panic(*exception.BussinessException("Cannot read the file"))
		}

		if int64(len(fileBytes)) > allowUploadSizeGeneral {
			panic(*exception.BussinessException(fmt.Sprintf("Upload file '%s' too large", file.Filename)))
		}

		fileContents[file.Filename] = fileBytes
	}

	var wg sync.WaitGroup
	errChan := make(chan error, len(fileContents))
	filePaths := make(map[string]string)
	index := 0
	for fileName, fileContent := range fileContents {
		wg.Add(1)
		go func(name string, content []byte, filePaths *map[string]string) {
			defer wg.Done()
			uuiDrandomStr := uuid.New().String() + "-" + utils.GenerateRandomString(10)
			replaceSpaceAndComma := strings.ReplaceAll(strings.ReplaceAll(name, " ", "_"), ",", "_")
			fileName := fmt.Sprintf("%d-%s", index, replaceSpaceAndComma)
			objectName := fmt.Sprintf("%s/%s", FileUploads, uuiDrandomStr)

			paths := *filePaths
			paths[fileName] = objectName

			log.Printf("Starting upload for file: %s", objectName)
			err := singleton.StoreFileBuff(content, objectName)
			if err != nil {
				log.Printf("Error uploading file %s: %v", name, err)
				errChan <- err
			} else {
				log.Printf("Successfully uploaded file: %s", name)
			}
			index++
		}(fileName, fileContent, &filePaths)
	}

	wg.Wait()
	close(errChan)

	if len(errChan) > 0 {
		panic(*exception.ServerErrorException("Failed to upload one or more files"))
	} else {
		log.Println("All files uploaded successfully")
	}

	return filePaths, nil
}

// Delete implements domain.FileUsecase.
func (f *FileUsecase) Delete(httpContext *gin.Context, filePaths []string) error {
	var wg sync.WaitGroup
	errChan := make(chan error, len(filePaths))

	for _, path := range filePaths {
		wg.Add(1)
		go func(filePath string) {
			defer wg.Done()
			err := singleton.DeleteFile(filePath)
			if err != nil {
				errChan <- err
			}
		}(path)
	}

	wg.Wait()
	close(errChan)

	// Collect errors
	var errs []error
	for err := range errChan {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		// Combine errors into a single error message
		errorMsgs := make([]string, len(errs))
		for i, err := range errs {
			errorMsgs[i] = err.Error()
		}
		return fmt.Errorf("failed to delete some files: %s", strings.Join(errorMsgs, "; "))
	}

	return nil
}

// Get implements domain.FileUsecase.
func (f *FileUsecase) Get(httpContext *gin.Context, filePath string) ([]byte, error) {
	return singleton.ReadFile(filePath)
}

// GetPresignUrlForUpload implements domain.FileUsecase.
func (f *FileUsecase) GetPresignUrlForUpload(httpContext *gin.Context, filePath string) (string, error) {
	targetPath := buildPathFileUpload(filePath)
	presignUrl, err := singleton.GetPresignURLUpload(targetPath)
	if err != nil {
		panic(*exception.ServerErrorException(fmt.Sprintf("Error when generate presign url : %s", err.Error())))
	}

	return presignUrl, nil
}

// GetPresignUrlForUpload implements domain.FileUsecase.
func (f *FileUsecase) GetPresignUrlForDownload(httpContext *gin.Context, filePath string) (string, error) {
	targetPath := buildPathFileUpload(filePath)
	presignUrl, err := singleton.GetPresignURLDownload(targetPath)
	if err != nil {
		panic(*exception.ServerErrorException(fmt.Sprintf("Error when generate presign url : %s", err.Error())))
	}

	return presignUrl, nil
}

func buildPathFileUpload(filePath string) string {
	uuiDrandomStr := uuid.New().String() + "-" + utils.GenerateRandomString(10)
	return fmt.Sprintf("%s/%s-%s", FileUploads, uuiDrandomStr, filePath)
}
