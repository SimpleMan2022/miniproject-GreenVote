package helpers

import (
	"errors"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func UploadUserImage(image *multipart.FileHeader) (string, error) {

	if _, err := os.Stat("public/images/users"); os.IsNotExist(err) {
		if err := os.MkdirAll("public/images/users", 0755); err != nil {
			return "", err
		}
	}
	fileName := uuid.New().String() + filepath.Ext(image.Filename)
	dst := filepath.Join("public/images/users", filepath.Base(fileName))
	outFile, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer outFile.Close()
	src, err := image.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	if _, err := io.Copy(outFile, src); err != nil {
		return "", err
	}
	return fileName, nil
}

func DeleteImageUser(image *string) error {
	imagePath := filepath.Join("public/images/users", *image)
	if err := os.Remove(imagePath); err != nil {
		return errors.New("image not found in directory")
	}
	return nil
}