package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sync"
	"time"

	"github.com/getsentry/sentry-go"
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

// Write implements the io.Writer interface. It processes and writes log messages to both files and standard output.
// The method categorizes logs based on content within square brackets (e.g., [category]) and writes them to separate log files.
//
// Parameters:
//   - content: byte slice containing the log message to be written
//
// Returns:
//   - n: number of bytes written
//   - err: any error encountered during writing
//
// The function performs the following operations:
//  1. Extracts log category from the message (defaults to "default" if not found)
//  2. Writes to category-specific log files when log driver is set to "file"
//  3. Always writes to standard output regardless of log driver setting
//
// Thread-safety is ensured through mutex locking.
func (logWriter *LogWriter) Write(content []byte) (n int, err error) {
	logMessage := string(content)

	category := "default"
	re := regexp.MustCompile(`\[(.*?)\]`)
	matches := re.FindStringSubmatch(logMessage)

	if len(matches) > 1 {
		category = matches[1]
	}

	sentry.AddBreadcrumb(&sentry.Breadcrumb{
		Category: "log",
		Message:  logMessage,
		Level:    sentry.LevelInfo,
		Data: map[string]interface{}{
			"timestamp": time.Now().Unix(),
		},
	})

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
