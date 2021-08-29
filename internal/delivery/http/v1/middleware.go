package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	formFileKey = "image"
)

var imageTypes = map[string]struct{}{
	"image/jpeg": {},
	"image/jpg":  {},
	"image/png":  {},
}

func (h *Handler) getImageFromMultipartFormData(c *gin.Context) ([]byte, string, error) {
	file, fileHeader, err := c.Request.FormFile(formFileKey)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	fileBody := make([]byte, fileHeader.Size)
	_, err = file.Read(fileBody)
	if err != nil {
		return nil, "", err
	}
	fileType := http.DetectContentType(fileBody)

	// Validating file type
	if _, ok := imageTypes[fileType]; !ok {
		return nil, "", errors.New("file type isn't supported")
	}

	return fileBody, fileType, nil
}
