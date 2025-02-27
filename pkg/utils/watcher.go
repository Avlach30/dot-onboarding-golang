package utils

import (
	"context"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

var (
	debounce = make(chan bool, 1)
)

func StartWatcher(cancel context.CancelFunc) {
	log.Println("Starting watcher...")

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				// Handle file operations for .go files
				if isGoFile(event.Name) {
					switch {
					case event.Op&fsnotify.Write == fsnotify.Write:
						cancel()
					case event.Op&fsnotify.Create == fsnotify.Create:
						cancel()
					case event.Op&fsnotify.Remove == fsnotify.Remove:
						cancel()
					case event.Op&fsnotify.Rename == fsnotify.Rename:
						// For renamed files, we need to re-add the watch for the new location
						if newPath := findNewPath(event.Name); newPath != "" {
							cancel()
						}
					}

					time.Sleep(2 * time.Second)
					restartApp()
				}

				// Handle new directory creation
				if event.Op&fsnotify.Create == fsnotify.Create {
					fi, err := os.Stat(event.Name)
					if err == nil && fi.IsDir() {
						addDirectoryToWatcher(watcher, event.Name)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Error:", err)
			}
		}
	}()

	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Walk through all directories and add them to the watcher
	addDirectoryToWatcher(watcher, currentDir)

	// Wait forever
	<-make(chan struct{})
}

func restartApp() {
	log.Println("Restarting asd the application...")

	if debounce != nil {
		select {
		case debounce <- true:
		default:
		}
	}

	log.Println(os.Args)
	// Build and run the application
	cmd := exec.Command("go", "run", "main.go", "--watch")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		log.Printf("Error restarting application: %v\n", err)
		return
	}
}

func findNewPath(oldPath string) string {
	// Get the parent directory
	dir := filepath.Dir(oldPath)

	// Read all files in the directory
	files, err := os.ReadDir(dir)
	if err != nil {
		return ""
	}

	// Look for .go files that were recently modified
	threshold := time.Now().Add(-5 * time.Second)
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".go") {
			info, err := file.Info()
			if err != nil {
				continue
			}
			if info.ModTime().After(threshold) {
				return filepath.Join(dir, file.Name())
			}
		}
	}
	return ""
}

func addDirectoryToWatcher(watcher *fsnotify.Watcher, path string) {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip if it's not a directory
		if !info.IsDir() {
			return nil
		}

		// Skip vendor and hidden directories
		if info.Name() == "vendor" || strings.HasPrefix(info.Name(), ".") {
			return filepath.SkipDir
		}

		log.Printf("Watching directory: %s", path)
		return watcher.Add(path)
	})

	if err != nil {
		log.Printf("Error walking directory %s: %v", path, err)
	}
}

func isGoFile(path string) bool {
	return strings.HasSuffix(path, ".go") && !strings.Contains(path, "vendor/")
}
