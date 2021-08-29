package service

import (
	"github.com/denis-shcherbinin/gora-studio-test-task/internal/entity"
	imageManager "github.com/denis-shcherbinin/gora-studio-test-task/internal/image-manager"
	"github.com/denis-shcherbinin/gora-studio-test-task/internal/repository"
)

type Photos interface {
	Upload(input UploadInput) (entity.Photo, error)
	GetAll() ([]entity.Photo, error)
	DeleteById(id int64) error
}

type Services struct {
	Photos
}

type Deps struct {
	Repo         *repository.Repositories
	ImageManager imageManager.ImageManager
}

func NewServices(deps Deps) *Services {
	return &Services{
		NewPhotoService(deps.Repo.Photos, deps.ImageManager),
	}
}
