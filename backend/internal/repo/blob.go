package repo

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// SaveLocalFile saves an uploaded image to the local filesystem
func (repo *Repo) SaveLocalFile(file *multipart.FileHeader, storagePath string) (string, error) {
	// Create the directory if it doesn't exist
	if _, err := os.Stat(storagePath); os.IsNotExist(err) {
		err := os.MkdirAll(storagePath, 0755)
		if err != nil {
			return "", fmt.Errorf("could not create storage directory: %w", err)
		}
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Create a unique filename to prevent overwrites
	filename := filepath.Base(file.Filename)
	dstPath := filepath.Join(storagePath, filename)

	dst, err := os.Create(dstPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}

	// Return the relative URL path the frontend will use
	// For example: /uploads/image.jpg
	return "/uploads/" + filename, nil
}
