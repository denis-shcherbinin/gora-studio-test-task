package service

import (
	"errors"
	"github.com/denis-shcherbinin/gora-studio-test-task/internal/entity"
	imageManager "github.com/denis-shcherbinin/gora-studio-test-task/internal/image-manager"
	"github.com/denis-shcherbinin/gora-studio-test-task/internal/repository"
	"github.com/sirupsen/logrus"
)

type PhotosService struct {
	repo         repository.Photos
	imageManager imageManager.ImageManager
}

func NewPhotoService(repo repository.Photos, imageManager imageManager.ImageManager) *PhotosService {
	return &PhotosService{
		repo:         repo,
		imageManager: imageManager,
	}
}

type UploadInput struct {
	FileBody    []byte
	ContentType string
}

// Upload uploads a new photo to storage and repository.
// It returns entity.Photo.
func (ps *PhotosService) Upload(input UploadInput) (entity.Photo, error) {

	imageUrl, err := ps.imageManager.Upload(imageManager.UploadInput{
		FileBody: input.FileBody,
		FileType: input.ContentType,
	})
	if err != nil {
		return entity.Photo{}, err
	}

	id, err := ps.repo.Upload(imageUrl)
	if err != nil {
		logrus.Errorf("error occurred while uploading file info to db: %v", err)

		if err = ps.repo.DeleteById(id); err != nil {
			return entity.Photo{}, err
		}

		return entity.Photo{}, err
	}

	return entity.Photo{
		Id:  id,
		Url: imageUrl,
	}, nil
}

// GetAll invokes Repositories.GetAll()
// It returns a slice with entity.Photo.
func (ps *PhotosService) GetAll() ([]entity.Photo, error) {
	photos, err := ps.repo.GetAll()
	if err != nil {
		return nil, errors.New("error occurred while getting all photos")
	}
	return photos, nil
}

// DeleteById removes the image with the passed id.
func (ps *PhotosService) DeleteById(id int64) error {
	photo, err := ps.repo.GetById(id)
	if err != nil {
		return err
	}

	if err = ps.repo.DeleteById(id); err != nil {
		return errors.New("error occurred while deleting photo")
	}

	if err = ps.imageManager.Delete(photo.Url); err != nil {
		return err
	}

	return nil
}
