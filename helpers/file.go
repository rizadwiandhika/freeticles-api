package helpers

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"

	"github.com/rizadwiandhika/miniproject-backend-alterra/config"
)

func SaveFile(file *multipart.FileHeader, destination string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	folder := filepath.Dir(destination)
	if !DoesExist(folder) {
		if err = os.MkdirAll(folder, 0755); err != nil {
			return err
		}
	}

	// Destination
	dst, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}

func DoesExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func DeleteFile(destination string) error {
	if !DoesExist(destination) {
		return nil
	}

	go logFile(destination)
	err := os.Remove(destination)
	if err != nil {
		return err
	}
	return nil
}

func logFile(filename string) {
	logFile := path.Join(config.WORKING_DIR, "ignores", "delete-file.log")
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("ERROR IN LOGGING", err.Error())
		return
	}

	defer f.Close()

	if _, err = f.WriteString(filename + "\n"); err != nil {
		fmt.Println("ERROR IN LOGGING", err.Error())
		return
	}
}
