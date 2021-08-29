package image_manager

import "math/rand"

const (
	letters        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	fileNameLength = 24
)

type UploadInput struct {
	FileBody []byte
	FileType string
}

type ImageManager interface {
	Upload(input UploadInput) (string, error)
	Delete(imageUrl string) error
}

func generateFileName() string {
	b := make([]byte, fileNameLength)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
