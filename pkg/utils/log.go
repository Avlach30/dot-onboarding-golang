package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sync"
	"time"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/config"
)

type LogWriter struct {
	files map[string]*os.File
	mu    sync.Mutex
}

func NewLogWriter() *LogWriter {
	return &LogWriter{
		files: make(map[string]*os.File),
	}
}

func (logWriter *LogWriter) Write(content []byte) (n int, err error) {
	logMessage := string(content)

	category := "default"
	re := regexp.MustCompile(`\[(.*?)\]`)
	matches := re.FindStringSubmatch(logMessage)

	if len(matches) > 1 {
		category = matches[1]
	}

	logWriter.mu.Lock()
	defer logWriter.mu.Unlock()

	logDriver := config.LogDriver

	if logDriver == "file" {
		logFolderPath := config.LogFolderPath
		file, exists := logWriter.files[category]
		if !exists {
			currentDate := time.Now().Format("2006-01-02")
			fileName := fmt.Sprintf("%s - %s.log", currentDate, category)
			file, err = os.OpenFile(filepath.Join(logFolderPath, fileName), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				return 0, err
			}
			logWriter.files[category] = file
		}

		n, err = file.Write([]byte(logMessage))
		if err != nil {
			return n, err
		}
	}

	_, stdoutErr := os.Stdout.Write([]byte(logMessage))
	if stdoutErr != nil {
		return n, stdoutErr
	}

	return n, nil
}
