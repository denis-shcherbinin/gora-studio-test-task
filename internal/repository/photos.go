package repository

import (
	"database/sql"
	"github.com/denis-shcherbinin/gora-studio-test-task/internal/entity"
	"github.com/sirupsen/logrus"
)

type PhotosRepo struct {
	db *sql.DB
}

func NewPhotosRepo(db *sql.DB) *PhotosRepo {
	return &PhotosRepo{
		db: db,
	}
}

// Upload inserts a new record(image url) into a photos table
// It returns the id of the new line.
func (pr *PhotosRepo) Upload(imageUrl string) (int64, error) {
	insertQuery := `INSERT INTO photos(image_url) VALUES(?)`
	statement, err := pr.db.Prepare(insertQuery)
	if err != nil {
		return 0, err
	}

	res, err := statement.Exec(imageUrl)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

// GetAll collects all image photos records from photos table
// and fills a slice with entity.Photo
// It returns a slice with entity.Photo.
func (pr *PhotosRepo) GetAll() ([]entity.Photo, error) {
	rows, err := pr.db.Query("SELECT * FROM photos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var photos []entity.Photo

	var id int64
	var url string
	for rows.Next() {
		if err = rows.Scan(&id, &url); err != nil {
			logrus.Errorf("error occured while scanning photo from db: %v", err)
			continue
		}
		photo := entity.Photo{
			Id:  id,
			Url: url,
		}

		photos = append(photos, photo)
	}

	return photos, nil
}

// DeleteById removes the record with the passed id.
func (pr *PhotosRepo) DeleteById(id int64) error {
	deleteQuery := `DELETE FROM photos where id=?`
	statement, err := pr.db.Prepare(deleteQuery)
	if err != nil {
		return err
	}

	if _, err = statement.Exec(id); err != nil {
		return err
	}

	return nil
}

// GetById finds a photo by the passed id
// It returns entity.Photo.
func (pr *PhotosRepo) GetById(id int64) (entity.Photo, error) {
	var imageUrl string

	row := pr.db.QueryRow(`SELECT image_url FROM photos WHERE id=?`, id)
	if err := row.Scan(&imageUrl); err != nil {
		return entity.Photo{}, err
	}

	return entity.Photo{
		Id:  id,
		Url: imageUrl,
	}, nil
}
