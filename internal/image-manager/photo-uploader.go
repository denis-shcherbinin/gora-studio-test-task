package image_manager

import (
	"fmt"
	"os"
	"strings"
)

type PhotoUploader struct {
	imageDir     string
	statImageDir string
}

func NewPhotoUploader(imageDir, statImageDir string) *PhotoUploader {
	return &PhotoUploader{
		imageDir: imageDir,
		statImageDir: statImageDir,
	}
}

// Upload uploads an image to a [pu.imageDir] directory
// It returns image url to the uploaded image.
func (pu *PhotoUploader) Upload(input UploadInput) (string, error) {
	err := os.MkdirAll("./image", os.ModePerm)
	if err != nil {
		return "", err
	}

	fileName := generateFileName()
	fullFileName := fmt.Sprintf("%s.%s", fileName, strings.Split(input.FileType, "/")[1])

	// Saving the image
	fileOnDisk, err := os.Create(fmt.Sprintf("%s/%s", pu.imageDir, fullFileName))
	if err != nil {
		return "", err
	}
	defer fileOnDisk.Close()

	_, err = fileOnDisk.Write(input.FileBody)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", pu.statImageDir, fullFileName), nil
}

// Delete removes an image from a [pu.imageDir] directory.
func (pu *PhotoUploader) Delete(imageUrl string) error {
	return os.Remove(strings.Replace(imageUrl, pu.statImageDir, pu.imageDir, 1))
}
