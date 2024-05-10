package helpers

import (
	"context"
	"errors"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/google/uuid"
	"github.com/spf13/viper"
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

func DeleteImage(path string, image *string) error {
	imagePath := filepath.Join(path, *image)
	if err := os.Remove(imagePath); err != nil {
		return errors.New("image not found in directory")
	}
	return nil
}

func UploadPlaceImage(body io.Reader) (string, error) {
	cloudinaryUrl := viper.GetString("CLOUDINARY_URL")
	c, err := cloudinary.NewFromURL(cloudinaryUrl)
	if err != nil {
		return "", err
	}
	uploadResult, err := c.Upload.Upload(
		context.Background(),
		body,
		uploader.UploadParams{},
	)
	if err != nil {
		return "", err
	}

	return uploadResult.SecureURL, nil
}
