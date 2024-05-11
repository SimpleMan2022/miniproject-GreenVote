package helpers

import (
	"context"
	"fmt"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/spf13/viper"
	"mime/multipart"
	"path/filepath"
	"strings"
)

func DeleteImage(imageURL string) error {
	parts := strings.Split(imageURL, "/")
	publicIDWithExtension := parts[len(parts)-2] + "/" + parts[len(parts)-1]
	publicID := strings.TrimSuffix(publicIDWithExtension, filepath.Ext(publicIDWithExtension))

	cloudinaryUrl := viper.GetString("CLOUDINARY_URL")
	c, err := cloudinary.NewFromURL(cloudinaryUrl)
	if err != nil {
		return err
	}
	fmt.Println(publicID)
	_, err = c.Upload.Destroy(context.Background(), uploader.DestroyParams{
		PublicID: publicID,
	})
	if err != nil {
		return fmt.Errorf("failed to delete image with public ID '%s': %w", publicID, err)
	}

	return nil

}

func UploadImageToCloudinary(body interface{}, path string) (string, error) {
	cloudinaryUrl := viper.GetString("CLOUDINARY_URL")
	c, err := cloudinary.NewFromURL(cloudinaryUrl)
	if err != nil {
		return "", err
	}

	var uploadResult *uploader.UploadResult
	switch v := body.(type) {
	case multipart.FileHeader:
		file, err := v.Open()
		if err != nil {
			return "", err
		}
		defer file.Close()

		uploadResult, err = c.Upload.Upload(
			context.Background(),
			file,
			uploader.UploadParams{
				Folder: path,
			},
		)
		if err != nil {
			return "", err
		}
	default:
		uploadResult, err = c.Upload.Upload(
			context.Background(),
			body,
			uploader.UploadParams{
				Folder: path,
			},
		)
		if err != nil {
			return "", err
		}
	}

	return uploadResult.SecureURL, nil
}
