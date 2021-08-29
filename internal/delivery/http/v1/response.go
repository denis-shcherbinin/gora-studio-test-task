package v1

import (
	"github.com/denis-shcherbinin/gora-studio-test-task/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type response struct {
	Message string `json:"message"`
}

type photoUploadResponse struct {
	Id  int64  `json:"id"`
	Url string `json:"url"`
}

type photoGetAllResponse struct {
	Photos []entity.Photo `json:"photos"`
}

func newResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, response{
		Message: message,
	})
}
