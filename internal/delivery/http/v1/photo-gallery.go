package v1

import (
	"github.com/denis-shcherbinin/gora-studio-test-task/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) initPhotoGalleryRoutes(api *gin.RouterGroup) {
	photo := api.Group("/photo")
	{
		photo.POST("/upload", h.uploadPhoto)
		photo.GET("/all", h.getAllPhotos)
		photo.DELETE("/delete", h.deletePhoto)
	}
}

// @Summary Photo upload
// @Tags Photo Gallery
// @Description Photo upload
// @ModuleID uploadPhoto
// @Accept mpfd
// @Produce json
// @Param image formData file true "Image [jpeg/jpg/png], 32 MB maximum"
// @Success 201 {object} photoUploadResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /photo/upload [post]
func (h *Handler) uploadPhoto(c *gin.Context) {
	fileBody, fileType, err := h.getImageFromMultipartFormData(c)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	photo, err := h.services.Photos.Upload(service.UploadInput{
		FileBody: fileBody, ContentType: fileType,
	})
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, photoUploadResponse{
		Id:  photo.Id,
		Url: photo.Url,
	})
}

// @Summary Getting all photos
// @Tags Photo Gallery
// @Description Getting all photos
// @ModuleID getAllPhotos
// @Produce json
// @Success 200 {object} photoGetAllResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /photo/all [get]
func (h *Handler) getAllPhotos(c *gin.Context) {
	photos, err := h.services.Photos.GetAll()
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, photoGetAllResponse{
		Photos: photos,
	})
}

type deletePhotoInput struct {
	Id int64 `json:"id"`
}

// @Summary Deleting Photo By Id
// @Tags Photo Gallery
// @Description Deleting photo by id
// @ModuleID deletePhoto
// @Accept json
// @Produce json
// @Param input body deletePhotoInput true "photo id to delete"
// @Success 200 {string} ok
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /photo/delete [delete]
func (h *Handler) deletePhoto(c *gin.Context) {
	var input deletePhotoInput

	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Photos.DeleteById(input.Id); err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
