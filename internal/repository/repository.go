package repository

import (
	"database/sql"
	"github.com/denis-shcherbinin/gora-studio-test-task/internal/entity"
	"github.com/sirupsen/logrus"
)

type Photos interface {
	Upload(imageUrl string) (int64, error)
	GetAll() ([]entity.Photo, error)
	DeleteById(id int64) error
	GetById(id int64) (entity.Photo, error)
}

type Repositories struct {
	Photos
}

func NewRepositories(db *sql.DB) *Repositories {
	photos := NewPhotosRepo(db)
	createPhotosTable(photos.db)

	return &Repositories{
		Photos: photos,
	}
}

func createPhotosTable(db *sql.DB) {
	schemaSQL := `CREATE TABLE IF NOT EXISTS photos (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	image_url VARCHAR(64)
);`
	statement, err := db.Prepare(schemaSQL)
	if err != nil {
		logrus.Errorf("error occured while preparing photos table: %v", err)
	}
	_, err = statement.Exec()
	if err != nil {
		logrus.Errorf("error occured while creating photos table: %v", err)
	}
}
